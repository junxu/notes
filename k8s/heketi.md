# Kuberntest存储功能模块部署文档
## 1. 部署架构

Kubernetes -----> Heketi  ------> Glusterfs集群

Heketi是glusterfs集群的一个代理，提供restapi方式管理volume的功能

注意：
a. 事先准备好glusterfs集群
b. glusterfs 集群上要准备一些裸盘给Heketi使用

## 2. Heketi部署
### 2.1 准备heketi.json文档
通过configmap保证heketi.json文件

### 2.1.1 通过文件heketi.json创建configmap
    kubectl create cm heketi --from-file=/root/xj/heketi.json  --namespace=kube-system
下面是一个heketi.json的配置文件示例

``` json
{
  "_port_comment": "Heketi Server Port Number",
  "port": "8080",

  "_use_auth": "Enable JWT authorization. Please enable for deployment",
  "use_auth": false,

  "_jwt": "Private keys for access",
  "jwt": {
    "_admin": "Admin has access to all APIs",
    "admin": {
      "key": "My Secret"
    },
    "_user": "User only has access to /volumes endpoint",
    "user": {
      "key": "My Secret"
    }
  },

  "_glusterfs_comment": "GlusterFS Configuration",
  "glusterfs": {
    "_executor_comment": [
      "Execute plugin. Possible choices: mock, ssh",
      "mock: This setting is used for testing and development.",
      "      It will not send commands to any node.",
      "ssh:  This setting will notify Heketi to ssh to the nodes.",
      "      It will need the values in sshexec to be configured.",
      "kubernetes: Communicate with GlusterFS containers over",
      "            Kubernetes exec api."
    ],
    "executor": "ssh", //重要选择ssh 

    "_sshexec_comment": "SSH username and private key file information",
    "sshexec": {
      "keyfile": "/etc/heketi/ssh-key", //重要 
      "user": "root",
      "port": "22",
      "fstab": "/etc/fstab"
    },

    "_kubeexec_comment": "Kubernetes configuration",
    "kubeexec": {
      "host" :"https://kubernetes.host:8443",
      "cert" : "/path/to/crt.file",
      "insecure": false,
      "user": "kubernetes username",
      "password": "password for kubernetes user",
      "namespace": "OpenShift project or Kubernetes namespace",
      "fstab": "Optional: Specify fstab file on node.  Default is /etc/fstab"
    },

    "_db_comment": "Database file name",
    "db": "/var/lib/heketi/heketi.db",  //关注

    "_loglevel_comment": [
      "Set log level. Choices are:",
      "  none, critical, error, warning, info, debug",
      "Default is warning"
    ],
    "loglevel" : "debug"
  }
}
```
    
### 2.1.2 直接创建configmap，configmap如下：

``` yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: heketi
  namespace: kube-system
data:
  heketi.json: |
    {
      "_port_comment": "Heketi Server Port Number",
      "port": "8080",

      "_use_auth": "Enable JWT authorization. Please enable for deployment",
      "use_auth": false,

      "_jwt": "Private keys for access",
      "jwt": {
        "_admin": "Admin has access to all APIs",
        "admin": {
          "key": "My Secret"
        },
        "_user": "User only has access to /volumes endpoint",
        "user": {
          "key": "My Secret"
        }
      },

      "_glusterfs_comment": "GlusterFS Configuration",
      "glusterfs": {
        "_executor_comment": [
          "Execute plugin. Possible choices: mock, ssh",
          "mock: This setting is used for testing and development.",
          "      It will not send commands to any node.",
          "ssh:  This setting will notify Heketi to ssh to the nodes.",
          "      It will need the values in sshexec to be configured.",
          "kubernetes: Communicate with GlusterFS containers over",
          "            Kubernetes exec api."
        ],
        "executor": "ssh",

        "_sshexec_comment": "SSH username and private key file information",
        "sshexec": {
          "keyfile": "/etc/heketi/ssh-key",
          "user": "root",
          "port": "22",
          "fstab": "/etc/fstab"
        },

        "_kubeexec_comment": "Kubernetes configuration",
        "kubeexec": {
          "host" :"https://kubernetes.host:8443",
          "cert" : "/path/to/crt.file",
          "insecure": false,
          "user": "kubernetes username",
          "password": "password for kubernetes user",
          "namespace": "OpenShift project or Kubernetes namespace",
          "fstab": "Optional: Specify fstab file on node.  Default is /etc/fstab"
        },

        "_db_comment": "Database file name",
        "db": "/var/lib/heketi/heketi.db",

        "_loglevel_comment": [
          "Set log level. Choices are:",
          "  none, critical, error, warning, info, debug",
          "Default is warning"
        ],
        "loglevel" : "debug"
      }
    }
```

### 2.2 准备sshkey
sshkey是能够glusterfs集群节点sshkey公钥

从glusterfs集群的sshkey公钥文件创建secret

```kubectl create secret generic  heketi --from-file=ssh-key=/root/xj/ssh-key -n kube-system```

### 2.3 部署heketi

#### 2.3.1 部署heketi的service和deployment
命令如下：

    kubectl create -f heketi-deploy.yaml

heketi-deploy.yaml文件如下：


``` yaml
---
kind: Service
apiVersion: v1
metadata:
  name: heketi
  labels:
    glusterfs: heketi-service
    heketi: service
  annotations:
    description: Exposes Heketi Service
spec:
  selector:
    glusterfs: heketi-pod
  ports:
  - name: heketi
    port: 8080
    targetPort: 8080
---
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: heketi
  labels:
    glusterfs: heketi-deployment
    heketi: deployment
  annotations:
    description: Defines how to deploy Heketi
spec:
  replicas: 1
  template:
    metadata:
      name: heketi
      labels:
        glusterfs: heketi-pod
        heketi: pod
    spec:
      containers:
      - image: registry.paas/library/heketi:5 //使用社区的heketi v5版本
        imagePullPolicy: IfNotPresent
        name: heketi
        ports:
        - containerPort: 8080
        volumeMounts:
        - name: db
          mountPath: "/var/lib/heketi"
        - name: config
          mountPath: /etc/heketi/heketi.json
          subPath: heketi.json
        - name: ssh-key
          mountPath: /etc/heketi/ssh-key
          subPath: ssh-key
        readinessProbe:
          timeoutSeconds: 3
          initialDelaySeconds: 3
          httpGet:
            path: "/hello"
            port: 8080
        livenessProbe:
          timeoutSeconds: 3
          initialDelaySeconds: 30
          httpGet:
            path: "/hello"
            port: 8080
      nodeSelector:
        kubernetes.io/hostname: gdpaasn5
      volumes:
      - name: db
        hostPath:
          path: /var/lib/heketi
      - name: config
        configMap:
          name: heketi
      - name: ssh-key
        secret:
          secretName: heketi
```

               
获取cluster ip `kubectl get svc -n kube-system | grep heketi`，执行`curl cluster-ip:8080/hello`正常返回`Hello from Heketi`


#### 2.3.2 设置topoloy
第一步：进入到heketi的pod，执行下面命令`heketi-cli topology load -j topology.json`, topoloy.json文件如下：

``` json
{
    "clusters": [
        {
            "nodes": [
                {
                    "node": {
                        "hostnames": {
                            "manage": [
                                "192.168.1.19"  //glusterfs集群节点的IP地址
                            ],
                            "storage": [
                                "192.168.1.19"  //glusterfs集群节点的IP地址
                            ]
                        },
                        "zone": 1
                    },
                    "devices": [
                        "/dev/vdb"		//该glusterfs节点上可供Heketi使用的裸盘
                    ]
                },
                {
                    "node": {
                        "hostnames": {
                            "manage": [
                                "192.168.1.20"  //glusterfs集群节点的IP地址
                            ],
                            "storage": [
                                "192.168.1.20" //glusterfs集群节点的IP地址
                            ]
                        },
                        "zone": 2
                    },
                    "devices": [
                        "/dev/vdb"
                    ]
                },
                {
                    "node": {
                        "hostnames": {
                            "manage": [
                                "192.168.1.21"
                            ],
                            "storage": [
                                "192.168.1.21"
                            ]
                        },
                        "zone": 2
                    },
                    "devices": [
                        "/dev/vdb"
                    ]
                }
            ]
        }
    ]
}
```


#### 2.3.4 heketi数据高可用
之前的heketi部署将db的数据保存在host的目录上，使用node selector绑定到指定节点。
可以将db数据也存到gluster集群上
*步骤1： 事先在gluster集群创建好一个volume， `gluster volume create heketi-xj replica  2 192.168.1.20:/mnt/xj 192.168.1.21:/mnt/xj force`， `gluster volume start heketi-xj`， k8s集群的slave节点要装glusterfs-clinet

* 步骤2: 创建endpoint和service，`kubectl create -f ep-gluster.yaml`， ep-gluster.yaml文件如下：

```
---
apiVersion: v1
kind: Endpoints
metadata:
 name: glusterfs-cluster
subsets:
 - addresses:
   - ip: 192.168.1.20
   ports:
   - port: 1
     protocol: TCP
 - addresses:
   - ip: 192.168.1.21
   ports:
   - port: 1
     protocol: TCP
 - addresses:
   - ip: 192.168.1.19
   ports:
   - port: 1
     protocol: TCP
---

kind: Service
apiVersion: v1
metadata:
  name: glusterfs-cluster
spec:
  ports:
  - port: 1
```
* 步骤3：执行2.3.1步骤，但yaml文件更改如下： 

``` yaml
---
kind: Service
apiVersion: v1
metadata:
  name: heketi
  labels:
    glusterfs: heketi-service
    heketi: service
  annotations:
    description: Exposes Heketi Service
spec:
  selector:
    glusterfs: heketi-pod
  ports:
  - name: heketi
    port: 8080
    targetPort: 8080
---
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: heketi
  labels:
    glusterfs: heketi-deployment
    heketi: deployment
  annotations:
    description: Defines how to deploy Heketi
spec:
  replicas: 1
  template:
    metadata:
      name: heketi
      labels:
        glusterfs: heketi-pod
        heketi: pod
    spec:
      containers:
      - image: registry.paas/library/heketi:5
        imagePullPolicy: IfNotPresent
        name: heketi
        ports:
        - containerPort: 8080
        volumeMounts:
        - name: db
          mountPath: "/var/lib/heketi"
        - name: config
          mountPath: /etc/heketi/heketi.json
          subPath: heketi.json
        - name: ssh-key
          mountPath: /etc/heketi/ssh-key
          subPath: ssh-key
        readinessProbe:
          timeoutSeconds: 3
          initialDelaySeconds: 3
          httpGet:
            path: "/hello"
            port: 8080
        livenessProbe:
          timeoutSeconds: 3
          initialDelaySeconds: 30
          httpGet:
            path: "/hello"
            port: 8080
      nodeSelector:
        kubernetes.io/hostname: gdpaasn5
      volumes:
      - name: db
        glusterfs:
          endpoints: glusterfs-cluster
          path: heketi-xj  //事先创建好的gluster volume
      - name: config
        configMap:
          name: heketi
      - name: ssh-key
        secret:
          secretName: heketi
```


## 3. K8S 对接
### 3.1 创建storageClass
执行命令`kubectl create -f sc.yaml`， sc.yaml文件如下：
```
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: glusterfs01
parameters:
  gidMax: "50000"
  gidMin: "40000"
  restauthenabled: "false"
  resturl: http://192.168.1.19:8080
  volumetype: replicate:3
provisioner: kubernetes.io/glusterfs
```

### 3.2 创建PVC
执行命令`kubectl create -f pvc.yaml`， pvc.yaml文件如下：
```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: test1
spec:
  accessModes:
  - ReadWriteMany
  resources:
    requests:
      storage: 10Gi
  storageClassName: glusterfs01
```


