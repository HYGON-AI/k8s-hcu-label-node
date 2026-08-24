# HCU Label Node

English | [简体中文](README.zh.md)

HCU Label Node is a Kubernetes DaemonSet component. Once deployed to the nodes of a cluster it automatically detects HCU hardware and drivers, and applies standardized labels to **nodes that contain an HCU**, so that [k8s-hcu-device-plugin](../k8s-hcu-device-plugin), the scheduler and other components can perform node filtering and resource discovery.

Current version: **v3.0.0**

Default image: `harbor.sourcefind.cn:5443/hcu/admin/base/hcu-label-node:v3.0.0`

## Features

| Capability | Description |
|------------|-------------|
| HCU hardware detection | Identifies vendor Display / Co-processor devices via `lspci` |
| Driver and runtime detection | Runs `/opt/hyhal/bin/hy-smi` to verify the driver is usable (5 second timeout) |
| Automatic labeling | Sets `hygon.com/hcu=true` on HCU nodes, and can additionally set model, compute unit count, VRAM and driver version labels |
| Label guarding | Watches the local node object and restores labels when they are deleted or modified; re-establishes the watch after it is interrupted |
| Silent on non-HCU nodes | When no HCU is present or the driver is not ready, the process idles instead of crash-looping |

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                       Kubernetes Cluster                        │
│  ┌───────────────────────────────────────────────────────────┐  │
│  │  hcu-label-node DaemonSet (one Pod per node)              │  │
│  │    ├─ lspci / hy-smi detection                            │  │
│  │    ├─ DCGM reads info of the first device                 │  │
│  │    └─ Update Node Labels ────────────┐                    │  │
│  └──────────────────────────────────────│────────────────────┘  │
│                                         ▼                       │
│  ┌───────────────────────────────────────────────────────────┐  │
│  │  Node (only HCU nodes receive labels)                     │  │
│  │    hygon.com/hcu=true                                     │  │
│  │    hygon.com/hcu.name=...            (optional)           │  │
│  │    hygon.com/hcu.cu-count=...        (optional)           │  │
│  │    hygon.com/hcu.vram=...G           (optional)           │  │
│  │    hygon.com/hcu.driver-version=...  (optional)           │  │
│  └───────────────────────────────────────────────────────────┘  │
│                          │                                      │
│                          ▼                                      │
│  ┌───────────────────────────────────────────────────────────┐  │
│  │  k8s-hcu-device-plugin / Scheduler / user workloads       │  │
│  │  select HCU nodes via nodeSelector / nodeAffinity         │  │
│  └───────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
```

## Node Labels

The label prefix is `hygon.com/hcu`.

| Label key | Example value | Enabled by default | Description |
|-----------|---------------|--------------------|-------------|
| `hygon.com/hcu` | `true` | **Yes** | Marks the node as having a usable HCU. Components such as the device-plugin DaemonSet rely on this label for node affinity |
| `hygon.com/hcu.name` | Device model string | Requires `-name` | Device name of the first card, reported by DCGM |
| `hygon.com/hcu.cu-count` | `120` | Requires `-cu-count` | Number of Compute Units on the first card |
| `hygon.com/hcu.vram` | `64G` | Requires `-vram` | Total VRAM of the first card, rounded up to GiB |
| `hygon.com/hcu.driver-version` | `6.3.27-V1.2.5` | Requires `-driver-version` | Taken from `hy-smi --showdriverversion`; the label is omitted if the read fails |

The default image entrypoint (see the `Dockerfile` CMD) already enables `-name`, `-cu-count`, `-vram` and `-driver-version`, so all of the labels above are written after deployment.

> **Note**: Device information is read from the **first physical card at index 0**. On heterogeneous multi-card nodes the labels only reflect the first card, so fine-grained scheduling should rely on the resources reported by the device-plugin or on node-level inspection.

## Relationship to Other Components

| Component | How it uses the labels produced here |
|-----------|--------------------------------------|
| [k8s-hcu-device-plugin](../k8s-hcu-device-plugin) | The DaemonSet uses `nodeAffinity` to schedule onto nodes labeled `hygon.com/hcu=true` |
| [k8s-hcu-scheduler](../k8s-hcu-scheduler) | Works with the device-plugin's node filtering in dynamic vHCU scenarios |
| [k8s-hcu-dra-driver](../k8s-hcu-dra-driver) | Uses its own `hcu=on` label as a nodeSelector and does **not** depend on this component; that label can be applied manually or alongside this component |

Deploying this component before the device-plugin is recommended, so that HCU nodes automatically receive `hygon.com/hcu=true`.

## Prerequisites

| Item | Description                                                                                  |
|------|----------------------------------------------------------------------------------------------|
| Kubernetes cluster | Ready, and able to schedule the DaemonSet onto the target nodes                              |
| HCU driver | The HCU driver is installed on the node and `/opt/hyhal/bin/hy-smi` is executable            |
| hyhal | The node has an `/opt/hyhal` directory (containing `bin/hy-smi` and the userspace libraries) |
| RBAC | The component needs `get` / `list` / `watch` / `update` permissions on the `nodes` resource  |
| Build environment (when building yourself) | A Linux host with CGO enabled and the hcu-dcgm build dependencies                            |

## Quick Start

### 1. Build the image (optional)

Run on a **Linux host**:

```bash
cd k8s-hcu-label-node
bash build.sh
```

The script will:

1. Run `go build` to produce the `hcu-label-node` binary
2. Build the image `harbor.sourcefind.cn:5443/hcu/admin/base/hcu-label-node:<git-tag>`
3. Export `hcu-label-node-<git-tag>.tar`

The image tag comes from the most recent git tag in the repository (`git describe --tags --abbrev=0`). If no tag exists, create one first or edit the image name in `build.sh`.

### 2. Deploy to the cluster

**Recommended: static manifest**

```bash
kubectl apply -f deployment/static/k8s-hcu-label-node.yaml
```

The manifest creates:

| Resource | Name | Description |
|----------|------|-------------|
| `ClusterRole` | `cr-node-labeller` | Read/write permissions on node labels |
| `ClusterRoleBinding` | `labeller` | Binds `node-labeller-sa` in the `kube-system` namespace |
| `ServiceAccount` | `node-labeller-sa` (`kube-system`) | Service account used by the component |
| `DaemonSet` | `hcu-label-node` (`kube-system`) | Runs one Pod on every node |

**Helm**

```bash
helm install hcu-label-node deployment/helm/hcu-label-node \
  -n kube-system \
  --set image.tag=v3.0.0
```

The Helm chart renders the same result as the static manifest (identical image, startup arguments, volume mounts and RBAC rules), and additionally allows the label switches, log level and scheduling policy to be adjusted through `values.yaml`:

| Setting | Default | Description |
|---------|---------|-------------|
| `namespace` | `kube-system` | Deployment namespace |
| `image.tag` | `v3.0.0` | Image tag |
| `labels.name` / `labels.cuCount` / `labels.vram` | `true` | Map to `-name` / `-cu-count` / `-vram` |
| `labels.driverVersion` | `true` | Maps to `-driver-version` |
| `log.verbosity` | `0` | Maps to `-v` |
| `rbac.create` / `serviceAccount.create` | `true` | Set to `false` if resources with the same names already exist in the cluster |

### 3. Verify

```bash
# DaemonSet readiness (every node should have one Pod; on non-HCU nodes the Pod is Running but idle inside)
kubectl -n kube-system get ds hcu-label-node
kubectl -n kube-system get pods -l name=hcu-label-node -o wide

# Inspect the labels of an HCU node (replace <node> with the real node name)
kubectl get node <node> --show-labels | tr ',' '\n' | grep hygon.com/hcu

# View the logs
kubectl -n kube-system logs -l name=hcu-label-node --tail=100
```

Expected logs on an HCU node:

```
HCU exists in this node
HCU driver loaded
```

Common log on a non-HCU node:

```
HCU not exists in this node, idle
```

## Runtime Behavior

The component runs through the following steps. If the hardware or driver check fails it **idles forever** (`select {}`), and the Pod stays Running without applying labels:

1. Detect vendor HCU devices with `lspci`; idle if none are present
2. Read the node name from the `NODE_NAME` environment variable
3. Pre-set `hygon.com/hcu=true` (before the driver is ready)
4. Run `/opt/hyhal/bin/hy-smi` to check whether the driver is loaded; idle on failure
5. Initialize DCGM with `dcgm.Init()`; idle on failure
6. Write the optional labels according to the startup flags (`name` / `cu-count` / `vram` / `driver-version`)
7. Watch the local node and rewrite the labels if they are modified; re-establish the watch after the channel closes

The node name is injected into the `NODE_NAME` environment variable by the DaemonSet through the Downward API (`spec.nodeName`). The process exits immediately if it is not injected.

All labels are submitted in a single `Update` request and retried automatically on write conflicts, which avoids mutual overwrites when other controllers modify the node object concurrently.

## Configuration Reference

### Startup flags

| Flag | Default | Description |
|------|---------|-------------|
| `-name` | `false` | Writes `hygon.com/hcu.name` |
| `-cu-count` | `false` | Writes `hygon.com/hcu.cu-count` |
| `-vram` | `false` | Writes `hygon.com/hcu.vram` |
| `-driver-version` | `false` | Writes `hygon.com/hcu.driver-version` |
| `-kubeconfig` | `/root/.kube/config` | Used when debugging from outside the cluster; falls back to InClusterConfig when the file does not exist |
| `-logtostderr` | — | glog writes to stderr (enabled by default in the image) |
| `-stderrthreshold` | — | Log level threshold (`INFO` by default in the image) |
| `-v` | — | glog verbosity level (`0` by default in the image) |

Default image CMD:

```
./hcu-label-node -name -cu-count -vram -driver-version -logtostderr=true -stderrthreshold=INFO -v=0
```

Both the static manifest and the Helm chart explicitly override `command` / `args`; removing that override falls back to the image entrypoint above.

### DaemonSet environment variables and mounts

| Setting | Description |
|---------|-------------|
| `NODE_NAME` | Comes from `spec.nodeName`, used to locate the node to label |
| `/dev` | Access to device nodes |
| `/usr/local` | Host runtime dependencies |
| `/opt` | hyhal userspace libraries and the `hy-smi` tool (read-only) |
| `privileged: true` | Required to access host devices and driver information |

## Example: Scheduling Pods with the Labels

Together with the device-plugin, `nodeSelector` can restrict workloads to HCU nodes:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: hcu-workload
spec:
  nodeSelector:
    hygon.com/hcu: "true"
  containers:
    - name: app
      image: your-hcu-image:tag
      resources:
        limits:
          hygon.com/hcu: 1
```

Filtering by card model (requires this component to run with `-name`):

```yaml
nodeSelector:
  hygon.com/hcu: "true"
  hygon.com/hcu.name: "BW3000"   # replace with the actual model
```

## Troubleshooting

### DaemonSet Pod in CrashLoopBackOff

- Check RBAC: whether the `ClusterRoleBinding` was created and whether the ServiceAccount has `nodes/update` permission
- Check the logs: `kubectl -n kube-system logs <pod-name>`

### An HCU node is not labeled with `hygon.com/hcu=true`

1. Confirm that `lspci | grep Co- || lspci | grep Display` produces output on the node
2. Confirm that `/opt/hyhal/bin/hy-smi` is executable and returns successfully
3. Confirm that the Pod mounts `/opt` and that `NODE_NAME` is injected
4. Check whether the Pod log stops at `DCGM init error` — this usually means the driver or hyhal version does not match

### Labels appear but are quickly overwritten or disappear

This component watches the node and restores labels after they are changed. If it conflicts with another controller, check whether an automation tool is rewriting node labels in bulk.

### Non-HCU nodes also run a label-node Pod

The DaemonSet has **no nodeSelector**, so it runs on every node. On non-HCU nodes the Pod stays Running and idle, which is expected behavior and does not affect unlabeled nodes.

### Labels on a multi-card node only reflect one card

This is by design: only `dcgm.GetDeviceInfo(0)` is read. For per-card information, use the resources registered by the device-plugin or node-level monitoring.

## Project Layout

```
k8s-hcu-label-node/
├── cmd/
│   └── main.go                      # Entrypoint: detection, labeling, watch loop
├── internal/pkg/labelnode/
│   └── labelnode.go                 # Kubernetes client and node label updates
├── internal/pkg/lib/                # hydmi / rocm-smi userspace libraries shipped with the image
├── deployment/
│   ├── static/
│   │   └── k8s-hcu-label-node.yaml  # Deployment manifest (RBAC + DaemonSet)
│   └── helm/hcu-label-node/         # Helm chart (equivalent to the static manifest)
├── .gitignore                       # Ignores dotfiles, build artifacts, images and archives
├── build.sh                         # Build script
├── Dockerfile
└── go.mod
```

## License

This project is licensed under the Apache License, Version 2.0.
Copyright 2026 Hygon Information Technology Co., Ltd.

See [LICENSE](LICENSE) and [NOTICE](NOTICE) for details.

Third-party Go module dependencies declared in `go.mod` are itemized in
[THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md), which records the repository, pinned version,
license and copyright notice of each dependency. All of them are permissive licenses
(Apache-2.0 / BSD-3-Clause / MIT / ISC) compatible with this project. That file is an index only:
the authoritative license and copyright text is the one shipped inside each dependency.
