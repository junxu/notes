
# 1. 有状态应用（statefulset）
## 1.1 列表
Request如下：

    
    GET  /api/v1/_raw/statefulset/namespace/{namespace}
Responce如下：

     {
       "kind": "StatefulSetList",
       "apiVersion": "apps/v1beta1",
       "metadata": {
        "selfLink": "/apis/apps/v1beta1/namespaces/cmss/statefulsets",
        "resourceVersion": "5712846"
       },
       "items": [
        {
         "metadata": {
          "name": "web4",
          "namespace": "cmss",
          "selfLink": "/apis/apps/v1beta1/namespaces/cmss/statefulsets/web4",
          "uid": "467c721a-92da-11e7-8976-fa163e5cad9c",
          "resourceVersion": "5420016",
          "generation": 1,
          "creationTimestamp": "2017-09-06T08:06:24Z",
          "labels": {
           "app": "nginx4"
          }
         },
         "spec": {
          "replicas": 3,
          "selector": {
           "matchLabels": {
            "app": "nginx4"
           }
          },
          "template": {
           "metadata": {
            "creationTimestamp": null,
            "labels": {
             "app": "nginx4"
            }
           },
           "spec": {
            "containers": [
             {
              "name": "nginx",
              "image": "registry.paas/library/nginx",
              "ports": [
               {
                "name": "web",
                "containerPort": 80,
                "protocol": "TCP"
               }
              ],
              "resources": {},
              "volumeMounts": [
               {
                "name": "www",
                "mountPath": "/usr/share/nginx/html"
               }
              ],
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File",
              "imagePullPolicy": "IfNotPresent"
             }
            ],
            "restartPolicy": "Always",
            "terminationGracePeriodSeconds": 10,
            "dnsPolicy": "ClusterFirst",
            "securityContext": {},
            "schedulerName": "default-scheduler"
           }
          },
          "volumeClaimTemplates": [
           {
            "metadata": {
             "name": "www",
             "creationTimestamp": null
            },
            "spec": {
             "accessModes": [
              "ReadWriteOnce"
             ],
             "resources": {
              "requests": {
               "storage": "2560Mi"
              }
             },
             "storageClassName": "glusterfs01"
            },
            "status": {
             "phase": "Pending"
            }
           }
          ],
          "serviceName": "nginx"
         },
         "status": {
          "observedGeneration": 1,
          "replicas": 3
         }
        }
       ]
      }


## 1.2 详情
Request如下：

      GET  /api/v1/_raw/statefulset/namespace/{namespace}/name/{app's name}
  
Responce如下：

     {
       "kind": "StatefulSet",
       "apiVersion": "apps/v1beta1",
       "metadata": {
        "name": "web4",
        "namespace": "cmss",
        "selfLink": "/apis/apps/v1beta1/namespaces/cmss/statefulsets/web4",
        "uid": "467c721a-92da-11e7-8976-fa163e5cad9c",
        "resourceVersion": "5420016",
        "generation": 1,
        "creationTimestamp": "2017-09-06T08:06:24Z",
        "labels": {
         "app": "nginx4"
        }
       },
       "spec": {
        "replicas": 3,
        "selector": {
         "matchLabels": {
          "app": "nginx4"
         }
        },
        "template": {
         "metadata": {
          "creationTimestamp": null,
          "labels": {
           "app": "nginx4"
          }
         },
         "spec": {
          "containers": [
           {
            "name": "nginx",
            "image": "registry.paas/library/nginx",
            "ports": [
             {
              "name": "web",
              "containerPort": 80,
              "protocol": "TCP"
             }
            ],
            "resources": {},
            "volumeMounts": [
             {
              "name": "www",
              "mountPath": "/usr/share/nginx/html"
             }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent"
           }
          ],
          "restartPolicy": "Always",
          "terminationGracePeriodSeconds": 10,
          "dnsPolicy": "ClusterFirst",
          "securityContext": {},
          "schedulerName": "default-scheduler"
         }
        },
        "volumeClaimTemplates": [
         {
          "metadata": {
           "name": "www",
           "creationTimestamp": null
          },
          "spec": {
           "accessModes": [
            "ReadWriteOnce"
           ],
           "resources": {
            "requests": {
             "storage": "2560Mi"
            }
           },
           "storageClassName": "glusterfs01"
          },
          "status": {
           "phase": "Pending"
          }
         }
        ],
        "serviceName": "nginx"
       },
       "status": {
        "observedGeneration": 1,
        "replicas": 3
       }
      }

## 1.3 编辑

    PUT api/v1/_raw/statefulset/namespace/{namespace}/name/{name}

     apiVersion: apps/v1beta1
     kind: StatefulSet
     metadata:
       name: web4
     spec:
       serviceName: "nginx"
       replicas: 3
       template:
         metadata:
           labels:
             app: nginx4
         spec:
           terminationGracePeriodSeconds: 10
           containers:
           - name: nginx
             image: registry.paas/library/nginx:1.12
             imagePullPolicy: IfNotPresent
             ports:
             - containerPort: 80
               name: web
             volumeMounts:
             - name: www
               mountPath: /usr/share/nginx/html
       volumeClaimTemplates:
       - metadata:
           name: www
         spec:
           accessModes: [ "ReadWriteOnce" ]
           storageClassName: glusterfs01
           resources:
             requests:
               storage: 2.5Gi   


## 1.4 删除

    DELETE /api/v1/_raw/statefulset/namespace/{namespace}/name/{name}


## 1.5 创建

     POST /api/v1/_raw/statefulset/namespace/{namespace}

     apiVersion: apps/v1beta1
     kind: StatefulSet
     metadata:
       name: web4
     spec:
       serviceName: "nginx"
       replicas: 3
       template:
         metadata:
           labels:
             app: nginx4
         spec:
           terminationGracePeriodSeconds: 10
           containers:
           - name: nginx
             image: registry.paas/library/nginx:1.12
             imagePullPolicy: IfNotPresent
             ports:
             - containerPort: 80
               name: web
             volumeMounts:   //重点
             - name: www
               mountPath: /usr/share/nginx/html
       volumeClaimTemplates:  //重点
       - metadata:
           name: www
         spec:
           accessModes: [ "ReadWriteOnce" ]
           storageClassName: glusterfs01
           resources:
             requests:
               storage: 2.5Gi


