/*
 * SPDX-License-Identifier: Apache-2.0
 * Copyright 2026 Hygon Information Technology Co., Ltd.
 */

package labelnode

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/golang/glog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
)

var kubeconfig = flag.String("kubeconfig", "/root/.kube/config", "absolute path to the kubeconfig file")

var (
	clientOnce sync.Once
	cachedErr  error
	cached     *kubernetes.Clientset
)

func buildConfig(kubeconfig string) (*rest.Config, error) {
	// Check if the kubeconfig file exists
	if _, err := os.Stat(kubeconfig); err == nil {
		// If the kubeconfig file exists, use it to build the configuration
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	// If the kubeconfig file does not exist, use InClusterConfig
	return rest.InClusterConfig()
}

// LabelNode sets labels on a node in a single Update call to reduce write conflicts.
func LabelNode(nodeName string, labels map[string]string) error {
	if nodeName == "" {
		return fmt.Errorf("node name is empty")
	}

	client, err := NewClient()
	if err != nil {
		return fmt.Errorf("get k8s client error: %v", err)
	}

	// The node object may be modified concurrently by other controllers, so re-read it on conflict
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		node, err := client.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
		if err != nil {
			return fmt.Errorf("get k8s node error: %v", err)
		}

		if node.Labels == nil {
			node.Labels = make(map[string]string, len(labels))
		}

		changed := false
		for k, v := range labels {
			if node.Labels[k] != v {
				node.Labels[k] = v
				changed = true
			}
		}
		if !changed {
			return nil
		}

		if _, err := client.CoreV1().Nodes().Update(context.TODO(), node, metav1.UpdateOptions{}); err != nil {
			return err
		}
		glog.V(5).Infof("Node %s labels updated: %v", nodeName, labels)
		return nil
	})
}

func NewClient() (*kubernetes.Clientset, error) {
	clientOnce.Do(func() {
		config, err := buildConfig(*kubeconfig)
		if err != nil {
			cachedErr = fmt.Errorf("build kube config error: %v", err)
			return
		}

		cached, err = kubernetes.NewForConfig(config)
		if err != nil {
			cachedErr = fmt.Errorf("get k8s client error: %v", err)
		}
	})
	return cached, cachedErr
}
