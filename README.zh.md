# HCU Label Node

[English](README.md) | 简体中文

HCU Label Node 是一个 Kubernetes DaemonSet 组件，部署到集群各节点后自动检测 HCU 硬件与驱动，并为 **含有 HCU 的节点** 打上标准化标签，供 [k8s-hcu-device-plugin](../k8s-hcu-device-plugin)、调度器及其他组件做节点筛选与资源发现。

当前版本：**v3.0.0**

默认镜像：`harbor.sourcefind.cn:5443/hcu/admin/base/hcu-label-node:v3.0.0`

## 功能特性

| 能力 | 说明 |
|------|------|
| HCU 硬件检测 | 通过 `lspci` 识别厂商的 Display / Co-processor 设备 |
| 驱动与运行时检测 | 执行 `/opt/hyhal/bin/hy-smi` 验证驱动可用（5 秒超时） |
| 自动打标 | 为 HCU 节点设置 `hygon.com/hcu=true`，并可附加型号、算力单元数、显存、驱动版本等标签 |
| 标签守护 | Watch 本节点对象，标签被误删或篡改时自动恢复；watch 中断后自动重建 |
| 非 HCU 节点静默 | 无 HCU 或驱动未就绪时进程进入空闲等待，不反复崩溃重启 |

## 架构概览

```
┌─────────────────────────────────────────────────────────────────┐
│                    Kubernetes Cluster                            │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  hcu-label-node DaemonSet（每个节点一个 Pod）              │   │
│  │    ├─ lspci / hy-smi 检测                                 │   │
│  │    ├─ DCGM 读取首张卡信息                                  │   │
│  │    └─ Update Node Labels ─────────────┐                   │   │
│  └───────────────────────────────────────│───────────────────┘   │
│                                          ▼                       │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  Node（仅 HCU 节点获得标签）                               │   │
│  │    hygon.com/hcu=true                                    │   │
│  │    hygon.com/hcu.name=...        （可选）                  │   │
│  │    hygon.com/hcu.cu-count=...    （可选）                  │   │
│  │    hygon.com/hcu.vram=...G       （可选）                  │   │
│  │    hygon.com/hcu.driver-version=...（可选）                │   │
│  └──────────────────────────────────────────────────────────┘   │
│                          │                                       │
│                          ▼                                       │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  k8s-hcu-device-plugin / Scheduler / 用户工作负载          │   │
│  │  通过 nodeSelector / nodeAffinity 选择 HCU 节点            │   │
│  └──────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
```

## 节点标签说明

标签前缀为 `hygon.com/hcu`。

| 标签键 | 示例值 | 是否默认启用 | 说明 |
|--------|--------|-------------|------|
| `hygon.com/hcu` | `true` | **是** | 标识该节点存在可用 HCU，device-plugin DaemonSet 等组件依赖此标签做节点亲和 |
| `hygon.com/hcu.name` | 设备型号字符串 | 需 `-name` | 来自 DCGM 首张卡的设备名称 |
| `hygon.com/hcu.cu-count` | `120` | 需 `-cu-count` | 首张卡的 Compute Unit 数量 |
| `hygon.com/hcu.vram` | `64G` | 需 `-vram` | 首张卡显存总量，向上取整到 GiB |
| `hygon.com/hcu.driver-version` | `6.3.27-V1.2.5` | 需 `-driver-version` | 取自 `hy-smi --showdriverversion`，读取失败时不写入该标签 |

镜像默认启动参数（见 `Dockerfile` CMD）已启用 `-name`、`-cu-count`、`-vram`、`-driver-version`，即部署后会写入全部上述标签。

> **说明**：设备信息取自 **索引为 0 的首张物理卡**。多卡异构节点上，标签仅反映首张卡规格，精细调度请结合 device-plugin 上报的资源或节点级巡检。

## 与其他组件的关系

| 组件 | 如何使用本组件输出的标签 |
|------|-------------------------|
| [k8s-hcu-device-plugin](../k8s-hcu-device-plugin) | DaemonSet 通过 `nodeAffinity` 调度到 `hygon.com/hcu=true` 的节点 |
| [k8s-hcu-scheduler](../k8s-hcu-scheduler) | 动态 vHCU 场景下，配合 device-plugin 的节点筛选 |
| [k8s-hcu-dra-driver](../k8s-hcu-dra-driver) | 使用独立标签 `hcu=on` 做 nodeSelector，**不依赖**本组件；可手动打标或与本组件并存 |

建议在部署 device-plugin 之前先部署本组件，确保 HCU 节点自动获得 `hygon.com/hcu=true`。

## 前置要求

| 项目 | 说明 |
|------|------|
| Kubernetes 集群 | 已就绪，且 DaemonSet 可调度到目标节点 |
| HCU 驱动 | 节点已安装 HCU 驱动，可执行 `/opt/hyhal/bin/hy-smi` |
| hyhal | 节点存在 `/opt/hyhal` 目录（含 `bin/hy-smi` 与用户态库） |
| RBAC | 组件需要对 `nodes` 资源的 `get` / `list` / `watch` / `update` 权限 |
| 构建环境（自行编译时） | Linux 宿主机，CGO 开启，具备 dcgm-dcu 编译依赖 |

## 快速开始

### 1. 构建镜像（可选）

在 **Linux 宿主机** 执行：

```bash
cd k8s-hcu-label-node
bash build.sh
```

脚本会：

1. `go build` 生成二进制 `hcu-label-node`
2. 构建镜像 `harbor.sourcefind.cn:5443/hcu/admin/base/hcu-label-node:<git-tag>`
3. 导出 `hcu-label-node-<git-tag>.tar`

镜像 tag 取自当前仓库最近 git tag（`git describe --tags --abbrev=0`）。无 tag 时需先打 tag，或手动修改 `build.sh` 中的镜像名。

### 2. 部署到集群

**推荐：静态清单**

```bash
kubectl apply -f deployment/static/k8s-hcu-label-node.yaml
```

该清单创建：

| 资源 | 名称 | 说明 |
|------|------|------|
| `ClusterRole` | `cr-node-labeller` | 节点标签读写权限 |
| `ClusterRoleBinding` | `labeller` | 绑定 `kube-system` 命名空间的 `node-labeller-sa` |
| `ServiceAccount` | `node-labeller-sa`（`kube-system`） | 组件使用的服务账号 |
| `DaemonSet` | `hcu-label-node`（`kube-system`） | 每个节点运行一个 Pod |

**Helm**

```bash
helm install hcu-label-node deployment/helm/hcu-label-node \
  -n kube-system \
  --set image.tag=v3.0.0
```

Helm Chart 与静态清单渲染结果一致（镜像、启动参数、挂载卷、RBAC 规则相同），
并额外支持通过 `values.yaml` 调整标签开关、日志级别与调度策略：

| 配置项 | 默认值 | 说明 |
|--------|--------|------|
| `namespace` | `kube-system` | 部署命名空间 |
| `image.tag` | `v3.0.0` | 镜像 tag |
| `labels.name` / `labels.cuCount` / `labels.vram` | `true` | 对应 `-name` / `-cu-count` / `-vram` |
| `labels.driverVersion` | `true` | 对应 `-driver-version` |
| `log.verbosity` | `0` | 对应 `-v` |
| `rbac.create` / `serviceAccount.create` | `true` | 集群已有同名资源时可置为 `false` |

### 3. 验证

```bash
# DaemonSet 就绪（每个节点应有一个 Pod，非 HCU 节点 Pod 为 Running 但内部空闲）
kubectl -n kube-system get ds hcu-label-node
kubectl -n kube-system get pods -l name=hcu-label-node -o wide

# 查看 HCU 节点标签（将 <node> 换成实际节点名）
kubectl get node <node> --show-labels | tr ',' '\n' | grep hygon.com/hcu

# 查看日志
kubectl -n kube-system logs -l name=hcu-label-node --tail=100
```

期望日志（HCU 节点）：

```
HCU exists in this node
HCU driver loaded
```

非 HCU 节点常见日志：

```
HCU not exists in this node, idle
```

## 运行逻辑

组件按以下顺序执行，硬件或驱动检测不通过则 **永久空闲**（`select {}`），Pod 保持 Running 但不打标：

1. `lspci` 检测是否存在厂商 HCU 设备；不存在则空闲
2. 从环境变量 `NODE_NAME` 读取节点名
3. 预先设置 `hygon.com/hcu=true`（驱动就绪前）
4. 执行 `/opt/hyhal/bin/hy-smi` 检测驱动是否加载；失败则空闲
5. `dcgm.Init()` 初始化 DCGM；失败则空闲
6. 根据启动参数写入可选标签（`name` / `cu-count` / `vram` / `driver-version`）
7. Watch 本节点，若标签被修改则重新写入；watch 通道关闭后自动重建

节点名由 DaemonSet 通过 Downward API（`spec.nodeName`）注入环境变量 `NODE_NAME`。未注入时进程直接退出。

所有标签在一次 `Update` 请求中提交，并在写冲突时自动重试，避免与其他控制器并发修改节点对象时相互覆盖。

## 配置参考

### 启动参数

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `-name` | `false` | 写入 `hygon.com/hcu.name` |
| `-cu-count` | `false` | 写入 `hygon.com/hcu.cu-count` |
| `-vram` | `false` | 写入 `hygon.com/hcu.vram` |
| `-driver-version` | `false` | 写入 `hygon.com/hcu.driver-version` |
| `-kubeconfig` | `/root/.kube/config` | 集群外调试时使用；文件不存在时自动走 InClusterConfig |
| `-logtostderr` | — | glog 输出到 stderr（镜像默认开启） |
| `-stderrthreshold` | — | 日志级别阈值（镜像默认 `INFO`） |
| `-v` | — | glog 详细级别（镜像默认 `0`） |

镜像默认 CMD：

```
./hcu-label-node -name -cu-count -vram -driver-version -logtostderr=true -stderrthreshold=INFO -v=0
```

静态清单与 Helm Chart 均显式覆盖 `command` / `args`；移除该覆盖时将使用上述镜像默认入口。

### DaemonSet 环境变量与挂载

| 配置项 | 说明 |
|--------|------|
| `NODE_NAME` | 来自 `spec.nodeName`，用于定位待打标节点 |
| `/dev` | 访问设备节点 |
| `/usr/local` | 宿主机运行时依赖 |
| `/opt` | hyhal 用户态库与 `hy-smi` 工具（只读） |
| `privileged: true` | 需要访问宿主机设备与驱动信息 |

## 使用标签调度 Pod 示例

配合 device-plugin 时，可通过 `nodeSelector` 将工作负载限制在 HCU 节点：

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

按卡型筛选（需本组件启用了 `-name`）：

```yaml
nodeSelector:
  hygon.com/hcu: "true"
  hygon.com/hcu.name: "BW3000"   # 替换为实际型号
```

## 故障排查

### DaemonSet Pod CrashLoopBackOff

- 检查 RBAC：`ClusterRoleBinding` 是否已创建，ServiceAccount 是否有 `nodes/update` 权限
- 查看日志：`kubectl -n kube-system logs <pod-name>`

### HCU 节点未打上 `hygon.com/hcu=true`

1. 确认节点上 `lspci | grep Co- || lspci | grep Display` 有输出
2. 确认 `/opt/hyhal/bin/hy-smi` 可执行且返回成功
3. 确认 Pod 已挂载 `/opt` 且 `NODE_NAME` 已注入
4. 查看 Pod 日志是否停在 `DCGM init error`——多为驱动或 hyhal 版本不匹配

### 标签存在但很快被覆盖/消失

本组件会 Watch 节点并在标签被改动后恢复。若与其他控制器冲突，检查是否有自动化工具在批量重写节点标签。

### 非 HCU 节点也有 label-node Pod

DaemonSet **无 nodeSelector**，会在所有节点运行。非 HCU 节点上 Pod 保持 Running 并空闲，属于预期行为，不影响无标签节点。

### 多卡节点标签只反映一张卡

设计如此：仅读取 `dcgm.GetDeviceInfo(0)`。如需 per-card 信息，请使用 device-plugin 注册的资源或节点级监控。

## 项目结构

```
k8s-hcu-label-node/
├── cmd/
│   └── main.go                      # 入口：检测、打标、Watch 循环
├── internal/pkg/labelnode/
│   └── labelnode.go                 # Kubernetes 客户端与节点标签更新
├── internal/pkg/lib/                # 随镜像分发的 hydmi / rocm-smi 用户态库
├── deployment/
│   ├── static/
│   │   └── k8s-hcu-label-node.yaml  # 部署清单（RBAC + DaemonSet）
│   └── helm/hcu-label-node/         # Helm Chart（与静态清单等价）
├── .gitignore                       # 忽略点文件、编译产物、镜像与压缩包
├── build.sh                         # 构建脚本
├── Dockerfile
└── go.mod
```

## License

This project is licensed under the Apache License, Version 2.0.
Copyright 2026 Hygon Information Technology Co., Ltd.

See [LICENSE](LICENSE) and [NOTICE](NOTICE) for details.
