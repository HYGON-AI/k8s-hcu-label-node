本文档涉及的yaml文件可以从[社区资源中心](https://download.sourcefind.cn:65024/5/main/Kubernetes%E6%8F%92%E4%BB%B6)进行下载，镜像可以从[光源镜像仓库](https://developer.sourcefind.cn/servicelist)直接拉取。

HCU-Label-Node以容器的形式发布，并以 DaemonSet 的形式部署在 Kubernetes 集群上。该组件运行时，将检查节点上是否存在HCU，若节点存在HCU将自动创建节点标签来描述HCU信息。

## 1.1 安装部署

k8s集群以DaemonSet形式部署HCU-Label-Node，DaemonSet部署命令如下：

```
kubectl apply -f k8s-hcu-label-node.yaml
```

## 1.2 部署验证

用户可以通过查看HCU-Label-Node的Pod运行情况验证HCU-Label-Node是否部署成功：

```
kubectl get pod -n kube-system -owide | grep hcu-label-node
```

如下输出表示HCU-Label-Node已经成功部署：

```
[root@localhost ~]# kubectl get pod -n kube-system -owide | grep hcu-label-node
hcu-label-node-2l52m                       1/1     Running   0               16s    10.233.70.20      worker-234   <none>           <none>
hcu-label-node-6w4jq                       1/1     Running   0               16s    10.233.122.159    worker-183   <none>           <none>
hcu-label-node-bg7bm                       1/1     Running   0               16s    10.233.67.18      worker-129   <none>           <none>
hcu-label-node-l7k29                       1/1     Running   0               16s    10.233.105.47     master-171   <none>           <none>
hcu-label-node-w66gr                       1/1     Running   0               16s    10.233.114.97     worker-200   <none>           <none>
```

HCU-Label-Node组件运行时，将检查节点上HCU的状态，并在确认配置了HCU的节点上，创建以下标签来描述HCU信息。

表 1.1 HCU-Label-Node 标签说明

| 标签                   | 类型   | 描述                                                         |
| ---------------------- | ------ | ------------------------------------------------------------ |
| hygon.com/hcu          | string | 若节点存在HCU设备，hygon.com/hcu标签取值为“true”；否则hygon.com/hcu为前缀的标签均不自动生成。 |
| hygon.com/hcu.name     | string | HCU型号 如：hygon.com/hcu.name=K100_AI                       |
| hygon.com/hcu.cu-count | string | HCU的CU数量 如：hygon.com/hcu.cu-count=120                   |
| hygon.com/hcu.vram     | string | HCU的显存容量 如：hygon.com/hcu.vram=64G                     |

备注：HCU信息标签hygon.com/hcu.name、hygon.com/hcu.cu-count、hygon.com/hcu.vram默认会自动生成，用户可以通过yaml以启动参数的形式开启或关闭。

```yaml
command: [ "./hcu-label-node" ]
args: [ "-name", "-cu-count", "-vram", "-driver-version", "-logtostderr=true", "-stderrthreshold=INFO", "-v=0"]
```

用户可以通过查看HCU节点k8s描述信息确认HCU标签是否设置成功：

```
kubectl describe node <node-name>
```

如下所示节点存在以hygon.com/hcu为前缀的标签，表示HCU-Label-Node组件已经成功部署。

```
[root@localhost ~]# kubectl describe node worker-183
Name:               worker-183
Roles:              worker
Labels:             beta.kubernetes.io/arch=amd64
                    beta.kubernetes.io/os=linux
                    hygon.com/hcu=true
                    hygon.com/hcu.cu-count=120
                    hygon.com/hcu.name=K100_AI
                    hygon.com/hcu.vram=64G
                    kubernetes.io/arch=amd64
                    kubernetes.io/hostname=worker-183
                    kubernetes.io/os=linux
                    node-role.kubernetes.io/worker=
```

# 