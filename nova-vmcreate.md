# nova vm 创建流程

							  |-->nova-schedule
	nova-api->nova-conductor--
							  |-->nova-compute

     nova-api
     nova.api.openstack.compute.servers.Controller.create
     ->nova.compute.api.API.create
       ->nova.compute.api.API._create_instance
         ->nova.conductor.api.ComputeTaskAPI.build_instances
     	  ->nova.conductor.rpcapi.ComputeTaskAPI.build_instances
     		->rpc调用conductor server端的build_instances
     
     nova-conductor
     nova.conductor.manager.ConductorManager.build_instances
     ->nova.conductor.manager.ComputeTaskManager.build_instances
       ->nova.scheduler.client.SchedulerClient.select_destinations
       ->nova.compute.rpcapi.ComputeAPI.build_and_run_instance
       
      nova-compute
      nova.compute.manager.ComputeManager.build_and_run_instance
      ->nova.compute.manager.ComputeManager._do_build_and_run_instance
        ->nova.compute.manager.ComputeManager._build_and_run_instance
     	 ->nova.compute.manager.ComputeManager._build_resources
     	   ->nova.compute.manager.ComputeManager._build_networks_for_instance
             ->nova.compute.manager.ComputeManager._allocate_network
			   ->nova.compute.manager.ComputeManager._allocate_network_async（调用neutron api创建port）
     	   ->nova.compute.manager.ComputeManager._prep_block_device
     	 ->self.driver.spawn
