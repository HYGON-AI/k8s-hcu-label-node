/*
 * SPDX-License-Identifier: Apache-2.0
 * Copyright 2026 Hygon Information Technology Co., Ltd.
 */

package main

import (
	"bytes"
	"context"
	"flag"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/HYGON-AI/hcu-dcgm/v3/pkg/dcgm"
	"github.com/HYGON-AI/k8s-hcu-label-node/internal/pkg/labelnode"
	"github.com/golang/glog"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/watch"
)

var labelPrefix = "hygon.com/hcu"

func idleForever() {
	select {}
}

func checkHCUStatus() bool {
	cmd := exec.Command("bash", "-c", "lspci | grep Display | grep Chengdu || lspci | grep Co-processor | grep Chengdu")
	output, err := cmd.Output()
	return !(err != nil || len(strings.TrimSpace(string(output))) == 0)
}

func checkHCUDriver() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/opt/hyhal/bin/hy-smi")

	return cmd.Run() == nil
}

// GetHCUDriverVersion returns the HCU driver version, or an empty string on failure.
func GetHCUDriverVersion() string {
	cmd := exec.Command("/opt/hyhal/bin/hy-smi", "--showdriverversion")

	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	// Match "Driver Version: xxx"
	re := regexp.MustCompile(`Driver Version:\s*([^\s]+)`)

	match := re.FindSubmatch(output)
	if len(match) != 2 {
		return ""
	}

	return string(bytes.TrimSpace(match[1]))
}

func updateLabels(nodeName string, name, cuCount, vram, driverVersion bool) map[string]string {
	labels := map[string]string{labelPrefix: "true"}

	// Keep the base label even when device info fails, otherwise downstream components treat the node as having no HCU
	device, err := dcgm.GetDeviceInfo(0)
	if err != nil {
		glog.Errorf("GetDeviceInfo error: %v", err)
	} else {
		if name {
			labels[labelPrefix+".name"] = device.Name
		}
		if cuCount {
			labels[labelPrefix+".cu-count"] = strconv.Itoa(device.ComputeUnitCount)
		}
		if vram {
			labels[labelPrefix+".vram"] = strconv.Itoa(int(math.Ceil(float64(device.GlobalMemSize)/math.Pow(1024, 3)))) + "G"
		}
	}

	// The driver version comes from hy-smi and does not depend on DCGM
	if driverVersion {
		if version := GetHCUDriverVersion(); version != "" {
			labels[labelPrefix+".driver-version"] = version
		} else {
			glog.Warning("Get HCU driver version failed, skip driver-version label")
		}
	}

	if err := labelnode.LabelNode(nodeName, labels); err != nil {
		glog.Errorf("Label node error: %v", err)
	}

	return labels
}

func main() {
	var name = flag.Bool("name", false, "HCU name label is required")
	var cuCount = flag.Bool("cu-count", false, "HCU CU number label is required")
	var vram = flag.Bool("vram", false, "HCU vram label is required")
	var driverVersion = flag.Bool("driver-version", false, "HCU driver version label is required")
	flag.Parse()
	defer glog.Flush()

	if !checkHCUStatus() {
		glog.Infoln("HCU not exists in this node, idle")
		idleForever()
	}

	glog.Infoln("HCU exists in this node")

	nodeName := os.Getenv("NODE_NAME")
	if nodeName == "" {
		glog.Exitln("NODE_NAME is empty, inject it via Downward API (spec.nodeName)")
	}
	glog.Infoln("Get NodeName : ", nodeName)

	if err := labelnode.LabelNode(nodeName, map[string]string{labelPrefix: "true"}); err != nil {
		glog.Errorf("Label node error: %v", err)
	}

	if !checkHCUDriver() {
		glog.Infoln("HCU driver not loaded, idle")
		idleForever()
	}

	glog.Infoln("HCU driver loaded")
	if err := dcgm.Init(); err != nil {
		glog.Errorf("DCGM init error: %v, idle", err)
		idleForever()
	}
	defer dcgm.ShutDown()

	expectedLabels := updateLabels(nodeName, *name, *cuCount, *vram, *driverVersion)

	client, err := labelnode.NewClient()
	if err != nil {
		glog.Fatalf("Create clientset error: %v", err)
	}

	// A watch ends on timeout or API server restart, so it must be re-established to keep guarding the labels
	for {
		watcher, err := client.CoreV1().Nodes().Watch(
			context.TODO(),
			metav1.ListOptions{
				FieldSelector: fields.OneTermEqualSelector("metadata.name", nodeName).String(),
			},
		)
		if err != nil {
			glog.Errorf("Create watcher error: %v, retrying...", err)
			time.Sleep(5 * time.Second)
			continue
		}

		expectedLabels = watchNode(watcher, nodeName, expectedLabels, *name, *cuCount, *vram, *driverVersion)
		// Avoid busy looping when the watch closes immediately
		time.Sleep(time.Second)
	}
}

// watchNode consumes watch events until the channel closes and returns the latest expected labels.
func watchNode(watcher watch.Interface, nodeName string, expectedLabels map[string]string,
	name, cuCount, vram, driverVersion bool) map[string]string {
	defer watcher.Stop()

	for event := range watcher.ResultChan() {
		if event.Type != watch.Modified {
			continue
		}

		node, ok := event.Object.(*v1.Node)
		if !ok {
			glog.Warningf("Unexpected watch object type %T", event.Object)
			continue
		}

		for k, v := range expectedLabels {
			if node.Labels[k] != v {
				glog.Infof("Label %s changed (expected=%s, got=%s), resetting...", k, v, node.Labels[k])
				expectedLabels = updateLabels(nodeName, name, cuCount, vram, driverVersion)
				break
			}
		}
	}

	glog.Infoln("Watch channel closed, re-establishing...")
	return expectedLabels
}
