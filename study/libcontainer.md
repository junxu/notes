libcontainer
==================
-------------------
Overview
========
modules
========
- factory

a. factory.go定义了`Factory`接口

    type Factory interface {
        //创建一个container
		Create(id string, config *configs.Config) (Container, error)

        //load一个已存在的container
		Load(id string) (Container, error)

		StartInitialization(pipefd uintptr) error

		Type() string
    }

b. factory_linux.go定义了`LinuxFactory`结构是`Factory`接口的一个实现

    type LinuxFactory struct {
        //Root directory for the factory to store state.
        Root string

        // InitPath is the absolute path to the init binary.
        InitPath string

        // InitArgs are arguments for calling the init responsibilities for spawning
        // a container.
        InitArgs []string

        Validator validate.Validator

        // NewCgroupsManager returns an initialized cgroups manager for a single container.
        NewCgroupsManager func(config *configs.Cgroup, paths map[string]string) cgroups.Manager
    }

`func New(root string, options ...func(*LinuxFactory) error) (Factory, error)`会创建一个LinuxFactory.
`func (l *LinuxFactory) Create(id string, config *configs.Config) (Container, error)`在l.root目下
建立一个l.root/id目录，实例化一个linuxContainer结构.
`func (l *LinuxFactory) Type() string`返回"libcontainer".
`func (l *LinuxFactory) StartInitialization(pipefd uintptr) (err error)`这个函数比较复杂，之后分析.
`func (l *LinuxFactory) Load(id string) (Container, error)`冲l.root/id目录下读取state.json，解码成linuxContainer结构.

- container

a. container.go
该文件中定义了跨平台的一些container数据结构和接口
定义了Container的接口

     type Container interface {
             ID() string
     
             Status() (Status, error)
     
             State() (*State, error)
     
             Config() configs.Config
     
             // Returns the PIDs inside this container. The PIDs are in the namespace of the calling process.
             // Some of the returned PIDs may no longer refer to processes in the Container, unless
             // the Container state is PAUSED in which case every PID in the slice is valid.
             Processes() ([]int, error)
     
             Stats() (*Stats, error)
     
             // Set cgroup resources of container as configured
             // We can use this to change resources when containers are running.
             Set() error
     
             // Start a process inside the container. Returns error if process fails to
             // start. You can track process lifecycle with passed Process structure.
             Start(process *Process) (err error)
     
             // Destroys the container after killing all running processes.
             // Any event registrations are removed before the container is destroyed.
             Destroy() error
     
             // If the Container state is RUNNING or PAUSING, sets the Container state to PAUSING and pauses
             // the execution of any user processes. Asynchronously, when the container finished being paused the
             // state is changed to PAUSED.
             // If the Container state is PAUSED, do nothing.
             Pause() error
     
             // If the Container state is PAUSED, resumes the execution of any user processes in the
             // Container before setting the Container state to RUNNING.
             // If the Container state is RUNNING, do nothing.
             Resume() error
     
             // NotifyOOM returns a read-only channel signaling when the container receives an OOM notification.
             NotifyOOM() (<-chan struct{}, error)
     } 

定义了State结构，表示一个运行状态的container的state

    type State struct {
            ID string `json:"id"`
    
            // InitProcessPid is the init process id in the parent namespace.
            InitProcessPid int `json:"init_process_pid"`
    
            // InitProcessStartTime is the init process start time.
            InitProcessStartTime string `json:"init_process_start"`
    
            // Path to all the cgroups setup for a container. Key is cgroup subsystem name
            // with the value as the path.
            CgroupPaths map[string]string `json:"cgroup_paths"`
    
            // NamespacePaths are filepaths to the container's namespaces. Key is the namespace type
            // with the value as the path.
            NamespacePaths map[configs.NamespaceType]string `json:"namespace_paths"`
    
            // Config is the container's configuration.
            Config configs.Config `json:"config"`
    }
b. container_linux.go

首先定义了linuxContainer结构，用于表示一个linux container，该结构实现了Container interface.

    type linuxContainer struct {
            id            string
            root          string
            config        *configs.Config
            cgroupManager cgroups.Manager
            initPath      string
            initArgs      []string
            initProcess   parentProcess
            m             sync.Mutex
    }

只分析一些主要接口的实现

`Pause, Resume，Destory, Processes, Stats` 都是通过调用cgroupManager的接口实现的。

`Stats`调用cgroupManager获取的是CgroupStats（包括cpu, memory, bikio的统计信息），然后网络相关的stats有时另外的实现），然后网络相关的stats有时另外的实现.

`Destory`实现，首先要判断是否配置了NEWPID namespace。如果没有配置需要做些额外的工作，需要先吧cgroup frozen，然后获取所有的pids
，对kill每个进程，然后Thawed cgroup。最后调用cgroupManager的Destroy接口。




- process
- init_linux
- rootfs
- configs
- cgroups
- nsenter
- nsinit






