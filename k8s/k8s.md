#### WSL2 启用systemd [原文](https://blog.csdn.net/hiqiming/article/details/125111806?spm=1001.2014.3001.5501)

##### 安装方法一

1. 确保默认的WSL本版为2

   ```
   wsl --set-default-version 2
   ```

2. 下载并解压缩 [distrod_wsl_launcher](https://github.com/nullpo-head/wsl-distrod/releases/latest/download/distrod_wsl_launcher-x86_64.zip)，解压提取exe文件

3. 按照提示安装自己需要的发行版本

4. （可选）若需要发行版本在Windows开机时启动，请在WSL中以下命令

   ```
   sudo /opt/distrod/bin/distrod enable --start-on-windows-boot
   ```

##### 安装方法二

1. 在wsl发行版中下载并运行最新的安装程序脚本

   ```
   curl -L -O "https://raw.githubusercontent.com/nullpo-head/wsl-distrod/main/install.sh"
   chmod +x install.sh
   sudo ./install.sh install
   ```

2. 启用wsl-distrod 随windows开机启动

   ```
   /opt/distrod/bin/distrod enable --start-on-windows-boot
   否则
   /opt/distrod/bin/distrod enable
   ```

3. 重新启动你的发行版，关闭wsl，在powershell中执行

   ```
   wsl --shutdown
   ```

##### 实际我是用这个方法解决的 [原文](https://blog.csdn.net/jiexijihe945/article/details/127490773)

1. 在/etc/下面增加一个wsl.conf文件，这个操作需要sudo权限，文件里面输入下面的内容：

   ```
   [boot]
   systemd=true
   ```

2. **切记一个字都不要错，否则可能导致wsl进不去，所以备份很重要**

3. 保存退出，并执行wsl --shutdown，重新进入wsl



#### WSL2安装K8S [原文](https://blog.csdn.net/qxxhjy/article/details/121863472)

1. ##### **关闭swap**

   临时关闭

   ```
   swapoff -a
   ```

   永久关闭

   ```
   1.切换到：C:\Users\【你的用户名】
   2.新建一个.wslconfig的配置文件
   3.添加下面这写配置
   [wsl2]
   swap=0 # 关闭swap
   
   [network]
   generateResolvConf = false # 解决域名解析失败的问题
   
   命令行：wsl --shutdown  关闭所有的虚拟机
   ```

   查看是否关闭了

   ```
   free
   ```

2. ##### **安装kebuctl [原文](https://www.cnblogs.com/tianmingzh/articles/15861671.html)**

   1. 更新apt软件包索引并安装使用Kubernetes apt存储库所需的软件包(linux系统已经换源):

      ```
      sudo apt-get update
      sudo apt-get install -y apt-transport-https ca-certificates curl
      ```

   2. 使用阿里云的镜像

      ```
      curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -
      ```

   3. 添加Kubernetes apt存储库:

      ```
      cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
      deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
      EOF
      ```

   4. 使用新存储库更新apt软件包索引并安装kubectl：

      ```
      sudo apt-get update
      sudo apt-get install -y kubectl
      ```

   5. 可以使用 `kubectl version`验证是否安装成功

3. ##### **安装kind**

   ```
   wget https://github.com/kubernetes-sigs/kind/releases/download/0.2.1/kind-linux-amd64
   mv kind-linux-amd64 kind
   chmod +x kind
   mv kind /usr/local/bin
   这个版本比较老旧，使用go安装吧，把之前的卸载啦
   mv /usr/local/bin/kind /usr/local/bin/oldkind
   go install sigs.k8s.io/kind@v0.18.0
   ```

   **如果遇到超时问题，可能是golang的代理问题**

   1.首先开启go module

   ```
   go env -w GO111MODULE=on     // Windows  export GO111MODULE=on        // macOS 或 Linux
   ```

   2.配置goproxy:

   阿里配置：

   ```
   go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/       // Windows  
   export GOPROXY=https://mirrors.aliyun.com/goproxy/          // macOS 或 Linux
   ```

   七牛云配置：

   ```
   go env -w GOPROXY=https://goproxy.cn      // Windows  
   export GOPROXY=https://goproxy.cn         // macOS 或 Linux
   ```

   **安装kind之后如果遇到Command 'kind' not found, did you mean:问题**

   ```
   检查GOPATH，GOROOT两个环境变量
   export GETHPATH=/mnt/c/APP/go-ethereum-master/build/bin
   export GOROOT=/usr/local/go
   export GOPATH=/root/gowww
   export PATH=$PATH:$GETHPATH:$GOROOT:$GOPATH/bin
   ```

4. ##### **安装单节点k8s**

   ```
   root@DESKTOP-5QQL3P7:~# kind create cluster
   Creating cluster "kind" ...
    ✓ Ensuring node image (kindest/node:v1.13.4)  
    ✓ Preparing nodes   
    ✓ Creating kubeadm config   
    ✓ Starting control-plane  ️ 
   Cluster creation complete. You can now use the cluster with:
   
   export KUBECONFIG="$(kind get kubeconfig-path --name="kind")"
   kubectl cluster-info
   ```

3. ##### **安装多节点的k8s**

   编写多节点的配置文件kind-3nodes.yaml

   ```
   kind: Cluster
   apiVersion: kind.x-k8s.io/v1alpha4
   containerdConfigPatches:
     - |-
       [plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"]
         endpoint = ["https://aa25jngun.mirror.aliyuncs.com"]
   nodes:
     - role: control-plane
       kubeadmConfigPatches:
       - |
         kind: InitConfiguration
         nodeRegistration:
           kubeletExtraArgs:
             node-labels: "ingress-ready=true"
       extraPortMappings:
       - containerPort: 30000
         hostPort: 8000
         protocol: TCP
       - containerPort: 30001
         hostPort: 8080
         protocol: TCP
       - containerPort: 30002
         hostPort: 4443
         protocol: TCP
     - role: worker
     - role: worker
   ```

   安装

   ```
   kind create cluster --name 3 --config ./kind-3nodes.yaml
   ```

6. ##### **出现如下问题解决**

   第一个问题timeout

   ```
   ERRO[15:10:23] timed out waiting for docker to be ready on node kind-control-plane
   Error: failed to create cluster: timed out waiting for docker to be ready on node kind-control-plane
   ```

   ```
   如果使用了 --wait 120 之后仍然遇到了 "failed to create cluster: timed out waiting for docker to be ready on node kind-control-plane" 错误，那么可以尝试以下方法：
   
   确保在 WSL2 中已经安装了 Docker。您可以使用 docker --version 命令来验证 Docker 是否已经正确安装并运行。
   
   确认在 WSL2 中安装了必要的 Linux 内核更新。您可以使用 wsl --status 命令检查是否有可用的更新，使用 wsl --update 命令来安装更新。
   
   检查您的机器上是否存在其他 Kubernetes 工具或者其他容器环境（如 minikube、Docker Desktop 等），这些工具可能会影响 kind create cluster 命令的执行。如果存在这样的工具，可以尝试将其停止或卸载后再次运行 kind create cluster 命令。
   
   如果您的网络连接缓慢或不稳定，可以尝试连接到其他网络或使用 VPN 来改善连接质量。
   
   确认您的机器配置满足 Kubernetes 的最低要求。例如，Kubernetes 至少需要 2GB 的内存和 2 个 CPU 核心才能运行。
   
   如果上述方法都无法解决问题，您可以尝试使用其他 Kubernetes 工具（如 minikube）或在另一台机器上运行 kind create cluster 命令。
   ```

   卸载掉之前环境

   ```
   apt --purge remove kubeadm
   apt --purge remove kubelet
   ```

   第二个问题（我是因为版本过高）

   ```
   Client Version: version.Info{Major:"1", Minor:"27", GitVersion:"v1.27.1", GitCommit:"4c9411232e10168d7b050c49a1b59f6df9d7ea4b", GitTreeState:"clean", BuildDate:"2023-04-14T13:21:19Z", GoVersion:"go1.20.3", Compiler:"gc", Platform:"linux/amd64"}
   Kustomize Version: v5.0.1
   ```

   ```
   W0428 16:47:35.007624   75724 loader.go:222] Config not found: /etc/kubernetes/kubelet.conf
   ```

   出现这个问题可能是版本问题，你可以安装低版本(1.21.2)去解决 [原文](https://zhuanlan.zhihu.com/p/426227999)

   ```
   apt --purge remove kubectl
   curl -x YOUR_PROXY_SERVER:PORT -LO https://dl.k8s.io/v1.21.2/bin/linux/amd64/kubectl
   chmod 755 ./kubectl
   sudo mv ./kubectl /usr/local/bin/kubectl
   kubectl version --client
   ```

   或者复制配置到文件

   ```
   cat $KUBECONFIG
   复制
   vim /etc/kubernetes/kubelet.conf
   粘贴
   :wq
   ```

5. ##### 查看kind K8s 运行结构

   ```
   #查看集群
   kubectl cluster-info --context kind
   #查看node
   kubectl get nodes
   #查看kube-system空间内运行的pod
   kubectl get pods -n kube-system
   ```
   

#### 使用K8S

1. ##### 部署nginx 测试

   ```
   kubectl create deployment nginx --image=nginx
   
   kubectl expose deployment nginx --port=80 --type=NodePort
   
   kubectl get pod,svc
   ```

   可能遇到的问题pod一直**pending**，先查看问题原因，因为kind创建的时候没有暴露端口 [可以看这篇文章](https://blog.51cto.com/tansong/4850215)

   ```
   kubectl describe pod nginx-6f4998596c-cdxcm
   ```
   运行下列命令创建新的 k8s cluster

   k8s.yaml文件

   ```yaml
   kind: Cluster
   apiVersion: kind.x-k8s.io/v1alpha4
   nodes:
   - role: control-plane
     kubeadmConfigPatches:
     - |
       kind: InitConfiguration
       nodeRegistration:
         kubeletExtraArgs:
           node-labels: "ingress-ready=true"
     extraPortMappings:
     - containerPort: 80
       hostPort: 80
       protocol: TCP
     - containerPort: 443
       hostPort: 443
       protocol: TCP
     - containerPort: 30000
       hostPort: 30000
       protocol: TCP
     - containerPort: 30434
       hostPort: 30434
       protocol: TCP
   ```

   ```sh
   vim k8s.yaml
   kind create cluster --name tsk8s --config /k8s/kind/kind-k8s.yaml
   ```

   extraPortMappings：把 K8s 容器（相当于 K8s 所在的服务器）端口暴露出来，这里暴露了 80、443、30000
   node-labels：只允许 Ingress controller 运行在有"ingress-ready=true"标签的 node 上

   可能遇到的问题，版本问题导致的

   ```
   Error: error loading config: decoding failure: no kind "Cluster" is registered for version "kind.x-k8s.io/v1alpha4" in scheme "sigs.k8s.io/kind/pkg/cluster/config/encoding/scheme.go:34"
   ```

   **重新安装kind版本之后问题解决**

   **第二个问题明明已经是Running状态，然后浏览器访问不通是什么原因**

   ```
   因为我们的环境是kind,所以先查看docker是否映射了端口出来
   docker ps
   使用映射的端口启动nginx
   ```

   nginx.yaml

   ```yaml
   kind: Service
   apiVersion: v1
   metadata:
     name: httpd-svc
   spec:
     selector:
         app: httpd-app
     type: NodePort #1
     ports:
     - port: 80
       nodePort: 30000 #2
   ```

   ```
   kubectl apply -f nginx.yaml
   ```

   

   你可以使用 kubectl patch 命令更新 Service 对象来指定 NodePort 的绑定地址。以下是一个示例命令：

   ```
   kubectl patch service nginx -p '{"spec": {"ports": [{"name": "http", "nodePort": 30434, "port": 80, "protocol": "TCP", "targetPort": 80}]}}'
   ```

   

2. 