# nova-compute
## 1. 目录结构
     compute
     |--api.py
     |--arch.py
     |--build_results.py
     |--cells_api.py
     |--claims.py
     |--cpumodel.py
     |--flavors.py
     |--hv_type.py
     |--__init__.py
     |--instance_actions.py
     |--manager.py
     |--monitors
     |  |--cpu_monitor.py
     |  |--__init__.py
     |  |--virt
     |     |--cpu_monitor.py
     |     |--__init__.py
     |--power_state.py
     |--resources
     |--base.py
     |--__init__.py
     |--vcpu.py
     |--resource_tracker.py
     |--rpcapi.py
     |--stats.py
     |--task_states.py
     |--utils.py
     |--vm_mode.py
     |--vm_states.py

* vm_mode.py文件定义了几种虚拟化的模式，如HVM,xen,uml等
* vm_states.py定义了虚拟机的状态以及允许软重启和硬重启的vm状态，如active,building,pasused等等.
* hv_type.py定义了hypervisor的类型，如qemu,kvm,hypver等
* arch.py定义了host的cpu架构，如arm，x86，powerpc等
* cpumodel.py
* power_state.py定义了虚拟机的电源状态，running，shutdown,paused还是suspened等.这些状态是冲hyervisor driver获取的.
* instance_actions.py定义了Possible actions on an instance.如create，delete，reboot等
* build_results.py定义了build an instance的结果
* task_states.py定义了vm在处理action时候的各种状态，如在执行snapshot时候，有image_snapshot，image_snapshot_pending，image_pending_upload，image_uploading这些状态. 
* monitors定义了Resource monitor API specification. virt给出了一个libvirt的cpu监控的实现.
* resources定义了一个the interface used for compute resource plugins. 给出了一个vcpu资源的实现。
* stats.py用来记录compute node workload stats. 记录数据有num_vm_%vm_states, num_task_%task_states, num_os_type_%ostype, num_proj_%project_id.
* flavors.py定义了一些flavor的处理函数与nova flavor相关的api对应（如create, delete, get_flavor, get_all_flavors）. 这些函数都是通过调用nova.objects.flavor.Flavor类来出来，这个类会做flavor一些信息的数据库查询，写入.
* rpcapi.py定义了ComputeAPI（Client side of the compute rpc API）和SecurityGroupAPI（Client side of the security group rpc API）.

## nova-compute启动

nova.cmd.compute.main
->nova.service.Service.create
->nova.service.Service.start
  ->nova.compute.manager.ComputeManager.init_host()
  ->nova.compute.manager.ComputeManager.pre_start_hook()
  ->nova.compute.manager.ComputeManager.post_start_hook()
  ->nova.compute.manager.ComputeManager.periodic_tasks()

