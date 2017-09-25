
# 1. 存储类（stroageClass）
## 1.1 列表
Request如下：

    ##GET  /api/v1/storageclass
    GET  /api/v1/_raw/storageclass
Responce如下：

     "listMeta": {
      "totalItems": 1
     },
     "storageClasses": [
      {
       "objectMeta": {
        "name": "slow",
        "creationTimestamp": "2017-05-18T08:36:05Z"
       },
       "typeMeta": {
        "kind": "storageclass"
       },
       "provisioner": "kubernetes.io/glusterfs",    //对应类型
       "parameters": {
        "gidMax": "50000",
        "gidMin": "40000",
        "restauthenabled": "false",
        "resturl": "http://10.132.47.151:8080",     //对应URL
        "volumetype": "replicate:3"                 //对应卷类型
       }
      }
     ]
    }

## 1.2 详情
Request如下：

    ##GET  /api/v1/storageclass/{name}
    GET  /api/v1/_raw/storageclass/{name}
  
Responce如下：

    {
      "objectMeta": {
       "name": "slow",
       "creationTimestamp": "2017-05-18T08:36:05Z"
      },
      "typeMeta": {
       "kind": "storageclass"
      },
      "provisioner": "kubernetes.io/glusterfs",
      "parameters": {
       "gidMax": "50000",
       "gidMin": "40000",
       "restauthenabled": "false",
       "resturl": "http://10.132.47.151:8080",
       "volumetype": "replicate:3"
      }
     }

## 1.3 编辑

只能修改label

    PUT api/v1/_raw/storageclass/name/{name}




## 1.4 删除

    DELETE /api/v1/_raw/storageclass/name/{name}

## 1.5 创建

     POST /api/v1/_raw/storageclass/

     apiVersion: storage.k8s.io/v1
     kind: StorageClass
     metadata:
       name: fast
       labels:
          key1: aaa
          key2: bbb
          key3: ccc
          key4: ddd
     provisioner: kubernetes.io/glusterfs
     parameters:
       resturl: "http://10.132.47.151:8080"
       restauthenabled: "true"
       restuser： “admin”
       secretNamespace: "default"
       secretName: "aaa"
       gidMin: "40000"
       gidMax: "50000"
       volumetype: "replicate:3"

# 2. 动态卷（persistentVolumeClain）
## 2.1 列表
   GET /api/v1/_raw/persistentvolumeclaim/namespace/{namespace}


## 2.2 详情
     GET  /api/v1/_raw/persistentvolumeclaim/namespace/default/name/c2
    {
      "kind": "PersistentVolumeClaim",
      "apiVersion": "v1",
      "metadata": {
       "name": "c2",    //名称
       "namespace": "default",
       "selfLink": "/api/v1/namespaces/default/persistentvolumeclaims/c2",
       "uid": "c3f6f75b-5ae7-11e7-9d8e-fa163e1cb62a",
       "resourceVersion": "10695941",
       "creationTimestamp": "2017-06-27T03:21:52Z",
       "annotations": {
        "pv.kubernetes.io/bind-completed": "yes",
        "pv.kubernetes.io/bound-by-controller": "yes",
        "volume.beta.kubernetes.io/storage-provisioner": "kubernetes.io/glusterfs"
       }
      },
      "spec": {
       "accessModes": [
        "ReadWriteOnce"
       ],
       "resources": {
        "requests": {
         "storage": "1Gi"					//容量
        }
       },
       "volumeName": "pvc-c3f6f75b-5ae7-11e7-9d8e-fa163e1cb62a",  //卷
       "storageClassName": "slow"     //存储池
      },
      "status": {
       "phase": "Bound",             //状态
       "accessModes": [              //访问模式
        "ReadWriteOnce"
       ],
       "capacity": {
        "storage": "1Gi"
       }
      }
     }
     }


## 2.3 编辑
只能编辑label

    PUT /api/v1/_raw/persistentvolumeclaim/namespace/{namesapce}/name/{name}

## 2.4 删除

    DELETE /api/v1/_raw/persistentvolumeclaim/namespace/{namesapce}/name/{name}
  
## 2.5 创建
    POST  /api/v1/_raw/persistentvolumeclaim/namespace/{namespace}

    yaml文件
    
    kind: PersistentVolumeClaim
    apiVersion: v1
    metadata:
      name: myclaim    //名称
      labels:
        key1: aa
        key2: bb
    spec:
      accessModes:
        - ReadWriteOnce   //访问方式
      resources:
        requests:
          storage: 8Gi   //容量
      storageClassName: slow  //存储池
