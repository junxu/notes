
# 1. 批任务（Job）

## 1.1 详情
Request如下：

    
    GET  /api/v1/_raw/job/namespace/{namespace}/name/{name}
  
Responce如下：

    {
      "kind": "Job",
      "apiVersion": "batch/v1",
      "metadata": {
       "name": "sleep1",        //名称
       "namespace": "default",
       "selfLink": "/apis/batch/v1/namespaces/default/jobs/sleep1",
       "uid": "42cc78ce-45b3-11e7-ade9-fa163e1cb62a",
       "resourceVersion": "7496402",
       "creationTimestamp": "2017-05-31T03:43:08Z",
       "labels": {
        "controller-uid": "42cc78ce-45b3-11e7-ade9-fa163e1cb62a",
        "job-name": "sleep1"
       }
      },
      "spec": {
       "parallelism": 1,
       "completions": 1,          //实例数量
       "activeDeadlineSeconds": 60,
       "selector": {
        "matchLabels": {
         "controller-uid": "42cc78ce-45b3-11e7-ade9-fa163e1cb62a"
        }
       },
       "template": {
        "metadata": {
         "name": "sleep1",
         "creationTimestamp": null,
         "labels": {
          "controller-uid": "42cc78ce-45b3-11e7-ade9-fa163e1cb62a",
          "job-name": "sleep1"
         }
        },
        "spec": {
         "containers": [
          {
           "name": "sleep1",
           "image": "busybox",
           "command": [
            "sleep",
            "220"
           ],
           "resources": {},
           "terminationMessagePath": "/dev/termination-log",
           "terminationMessagePolicy": "File",
           "imagePullPolicy": "Always"
          }
         ],
         "restartPolicy": "Never",
         "terminationGracePeriodSeconds": 30,
         "dnsPolicy": "ClusterFirst",
         "securityContext": {},
         "schedulerName": "default-scheduler"
        }
       }
      },
      "status": {
       "conditions": [
        {
         "type": "Failed",
         "status": "True",
         "lastProbeTime": "2017-05-31T03:46:54Z",
         "lastTransitionTime": "2017-05-31T03:46:54Z",
         "reason": "DeadlineExceeded",
         "message": "Job was active longer than specified deadline"
        }
       ],
       "startTime": "2017-05-31T03:43:08Z",
       "succeeded": 1
      }
     }

## 1.3 编辑

无编辑
     #PUT  /api/v1/_raw/job/namespace/{namespace}/name/{name}

## 1.4 删除

    DELETE /api/v1/_raw/job/namespace/{namespace}/name/{name}

## 1.5 创建

     POST /api/v1/_raw/job/namespace/{namespace}

BODY示例：

    apiVersion: batch/v1
    kind: Job
    metadata:
      name: sleep1
    spec:
      completions: 3     //对应实例数
      template:
        spec:
          containers:
          - name: sleep1
            image: busybox
            imagePullPolicy: IfNotPresent
            command: [echo,  "220"]
          restartPolicy: Never

