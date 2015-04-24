# docker

---------------------------------

## 1. image目录

### 1.1 目录结构

    image
    |--graph.go
    |--image.go
    |--spec

* spec目录下是docker image的[Specification](https://github.com/docker/docker/blob/master/image/spec/v1.md).

- graph.go定义了`Graph interface`接口

	type Graph interface {
        Get(id string) (*Image, error)
        ImageRoot(id string) string
        Driver() graphdriver.Driver
	}

	>graph/graph.go中的`Graph struct`是对该接口的一个实现

+ image.go定义了`Image struct`结构

	     type Image struct {
             ID              string            `json:"id"`
             Parent          string            `json:"parent,omitempty"`
             Comment         string            `json:"comment,omitempty"`
             Created         time.Time         `json:"created"`
             Container       string            `json:"container,omitempty"`
             ContainerConfig runconfig.Config  `json:"container_config,omitempty"`
             DockerVersion   string            `json:"docker_version,omitempty"`
             Author          string            `json:"author,omitempty"`
             Config          *runconfig.Config `json:"config,omitempty"`
             Architecture    string            `json:"architecture,omitempty"`
             OS              string            `json:"os,omitempty"`
             Size            int64
     
             graph Graph
     	}



	>`image`的对外接口
	
		func (img *Image) SetGraph(graph Graph)
		//将img.graph=graph
		func (img *Image) SaveSize(root string) error
		//将img.Size存到“root/laysize”文件
		func (img *Image) SaveCheckSum(root, checksum string) error
		//将checksum写入到“root/checksum”文件
		func (img *Image) GetCheckSum(root string) (string, error)
		//从“root/checksum”文件读取校验值
		func (img *Image) RawJson() ([]byte, error)
		//step1： img.root()获取该img的root目录
		//step2： 读取该root目录下的json文件
		func (img *Image) TarLayer() (arch archive.Archive, err error)
		//调用img.graph.Driver().Diff接口
		func (img *Image) History() ([]*Image, error)
		//调用img.WalkHistory获取所有的祖先image信息
		func (img *Image) WalkHistory(handler func(*Image) error) (err error)
		//遍历祖先，主要使用img.GetParent来实现
		func (img *Image) GetParent() (*Image, error)
		//获取父亲的image信息
		//调用img.graph.Get(img.Parent)
		func (img *Image) root() (string, error)
		//获取image的root目录
		//img.graph.ImageRoot(img.ID)
		func (img *Image) GetParentsSize(size int64) int64
		//img.Size只记录本层的大小
		//该函数是获取该img所有祖先的Size之和
		func (img *Image) Depth() (int, error)
		//查看有多少祖先，即当前img是第多少层
		func (img *Image) CheckDepth() error
		//检查层级是否超过了MaxImageDepth(=127)
	

	> 定义了以下外部函数

		func LoadImage(root string) (*Image, error)
		//step1: 将“root/json”这个json格式的文件解析成image结构
		//step2: 读取“root/layersize”这个文件的内如，将其复制给image.Size,若该文件没有内容则将image.Size=-1
	
		func StoreImage(img *Image, layerData archive.ArchiveReader, root string) (err error)
		//stores file system layer data for the given image to the
		//image's registered storage driver. Image metadata is stored in a file
		//at the specified root directory.
		//step1: 如果layerData有值，则调用img.graph.Driver().ApplyDiff(img.ID, img.Parent, layerData)
		//step2: 调用img.SaveSize(root)将img.Size写入到“root/laysize”文件中
		//step3: 将img结构以jason的格式存入到“root/json”文件

		func NewImgJSON(src []byte) (*Image, error)
		//Build an Image object from raw json data
	
+ 补充

在/var/lib/docker目录下
会有个子目录graph,graph目录下对于每个image的layer都有一个子目录，目录名是image layer的id，再在这个目录下的是image.go写的内容。
* _**疑问**_

	>image结构和image.Size为什么分开存储？


## 2. graph目录

### 1.1 目录结构

    graph
     |--export.go
     |--graph.go
     |--history.go
     |--import.go
     |--list.go
     |--load.go
     |--load_unsupported.go
     |--manifest.go
     |--pull.go
     |--push.go
     |--service.go
     |--tag.go
     |--tags.go
     |--viz.go

其中最重要的文件graph.go，他是对image/graph.go中定义的`Graph interface`的实现.

其他文件是对应docker api在内部是处理函数.

### 1.2 graph.go
定义了`Graph struct`结构，该结构实现了image目录下的`Graph interface`接口.

     // A Graph is a store for versioned filesystem images and the relationship between them.
     type Graph struct {
             Root    string
             idIndex *truncindex.TruncIndex
             driver  graphdriver.Driver
     }

export出来的函数：

     func NewGraph(root string, driver graphdriver.Driver) (*Graph, error)
	 //创建一个Graph，输入的参数是graphdriver.Driver
     //step1: root是一个目录，如果该目录不存在将创建
     //step2: 创建Graph结构
     //step3： graph.restore()回复Graph(根据root目录下记录的数据做restore)
     func bufferToFile(f *os.File, src io.Reader) (int64, error)
     //把src的内容copy到f -> f.Sync() -> f.Seek(0, 0)
     func SetupInitLayer(initLayer string) error
     func isNotEmpty(err error) bool
     
     
     func (graph *Graph) restore() error
     //step1： 获取graph.root目录的下的文件,这些文件名即是image的id
	 //step2: 调用graph.driver.Exists(id)判断该image是否存在
     //step3： 将所有存在的ids用来初始化graph.idIndex

     func (graph *Graph) IsNotExist(err error)
     func (graph *Graph) Exists(id string) bool 
	 //returns true if an image is registered at the given id.
     //调用graph.Get(id)来判断

     func (graph *Graph) Create(layerData archive.ArchiveReader, containerID, containerImage, comment, author string, containerConfig, config *runconfig.Config) (*image.Image, error)
	 //creates a new image and registers it in the graph.
     //step1: 创建一个Image结构（使用传进来的参数）
     //step2: graph.Register(img, layerData)（注册image）
     func (graph *Graph) Register(img *image.Image, layerData archive.ArchiveReader) (err error)
	 //step1: 判断该image是否已经存在，graph.Exists(img.ID)
	 //step2: 清空graph.ImageRoot(img.ID)该目录
   	 //step3: 确保driver中也不存在该ID（以graph的信息为准），graph.driver.Remove(img.ID)
	 //step4: 创建一个临时目录，graph.Mktemp("")
	 //step5: Create root filesystem in the driver，graph.driver.Create(img.ID, img.Parent)
     //step6: 设置img的Graph成员，img.SetGraph(graph) 
	 //step7: Apply the diff/layer， image.StoreImage(img, layerData, tmp)
	 //step8: 将临时目录重命名成"root/img.ID"
     //step9: graph.idIndex.Add(img.ID)
     func (graph *Graph) TempLayerArchive(id string, sf *streamformatter.StreamFormatter, output io.Writer) (*archive.TempArchive, error)
     //creates a temporary archive of the given image's filesystem layer.
	 //主要调用的函数是image.TarLayer()
     func (graph *Graph) Mktemp(id string) (string, error)
	 //在graph.root目录生成一个目录"_tmp/random_id"
     func (graph *Graph) newTempFile() (*os.File, error)
	 //step1: 调用graph.Mktemp("")生成一个临时目录
     //step2: 在上述临时目录中随机生成一个临时文件
     func (graph *Graph) Delete(name string) error
	 //step1: 由name获取完整id， graph.idIndex.Get(name)
	 //step2: graph删除id, graph.idIndex.Delete(id)
     //step3: graph删除id, 将graph.ImageRoot(id)目录删除
	 //step4: driver删除id, graph.driver.Remove(id)
     func (graph *Graph) Map() (map[string]*image.Image, error)
	 //a list of all images in the graph, addressable by ID.
 	 //主要是调用graph.walkAll完成
     func (graph *Graph) walkAll(handler func(*image.Image)) error
	 //iterates over each image in the graph, and passes it to a handler.
	 //step1: 读取graph.Root目录，ioutil.ReadDir(graph.Root)
	 //step2: 对于每个目录项名称，调用graph.Get获取对应的Image结构
	 //step3: 将img结构传给handler来出来、
     func (graph *Graph) ByParent() (map[string][]*image.Image, error)
	 //returns a lookup table of images by their parent.
	 //If an image of id ID has 3 children images, then the value for key ID will be a list of 3 images.
	 //也是调用graph.walkAll来完成的
     func (graph *Graph) Heads() (map[string]*image.Image, error)
	 //returns all heads in the graph, keyed by id.
	 // A head is an image which is not the parent of another image in the graph.
 	 //主要是调用graph.walkAll完成
     
     func (graph *Graph) ImageRoot(id string) string
     //返回值是“graph.Root/id”路径名
     func (graph *Graph) Driver() graphdriver.Driver
     //返回graph.driver
     func (graph *Graph) Get(name string) (*image.Image, error)
	 //returns the image with the given id
	 //name应该是一个短id
     //step1: graph.idIndex.Get(name)获取name的完整id
     //step2: 从root/id目录下的json文件生成image结构，img = image.LoadImage(graph.ImageRoot(id))
     //step3: img.SetGraph(graph)
     //step4: 若img.Size < 0，则graph.driver.DiffSize(img.ID, img.Parent)获取size，并赋值给img.Size
     

### 1.3 other



## 1. api目录

### 1.1 目录结构

    api
    |--common.go
    |--server
    |  |--form.go  
    |  |--profiler.go  
    |  |--server.go  
    |  |--server_linux.go  
    |  |--tcp_socket.go  
    |  |--unix_socket.go
    |--client
    |  |--xxx.go
    |--types
    |  |--stats.go
    |  |--types.go

### 1.2 client目录
client目录下是docker client cli的实现.

### 1.3 types目录
stats.go文件定义了各种统计数据结构：
针对blkio，cpu，memory，network

### 1.4 server目录

#### 1.4.1 form.go

定义了两个函数，用于处理表单中的bool值和int值.

#### 1.4.2 

    "GET": {
            "/_ping":                          ping,
            "/events":                         getEvents,
            "/info":                           getInfo,					->eng.ServeHTTP
            "/version":                        getVersion,				->eng.ServeHTTP
            "/images/json":                    getImagesJSON,			
            "/images/viz":                     getImagesViz,			->eng.ServeHTTP
            "/images/search":                  getImagesSearch,			
            "/images/get":                     getImagesGet,
            "/images/{name:.*}/get":           getImagesGet,
            "/images/{name:.*}/history":       getImagesHistory,
            "/images/{name:.*}/json":          getImagesByName,
            "/containers/ps":                  getContainersJSON,
            "/containers/json":                getContainersJSON,
            "/containers/{name:.*}/export":    getContainersExport,
            "/containers/{name:.*}/changes":   getContainersChanges,
            "/containers/{name:.*}/json":      getContainersByName,
            "/containers/{name:.*}/top":       getContainersTop,
            "/containers/{name:.*}/logs":      getContainersLogs,
            "/containers/{name:.*}/stats":     getContainersStats,
            "/containers/{name:.*}/attach/ws": wsContainersAttach,
            "/exec/{id:.*}/json":              getExecByID,
    },
    "POST": {
            "/auth":                         postAuth,
            "/commit":                       postCommit,
            "/build":                        postBuild,
            "/images/create":                postImagesCreate,
            "/images/load":                  postImagesLoad,
            "/images/{name:.*}/push":        postImagesPush,
            "/images/{name:.*}/tag":         postImagesTag,
            "/containers/create":            postContainersCreate,
            "/containers/{name:.*}/kill":    postContainersKill,
            "/containers/{name:.*}/pause":   postContainersPause,
            "/containers/{name:.*}/unpause": postContainersUnpause,
            "/containers/{name:.*}/restart": postContainersRestart,
            "/containers/{name:.*}/start":   postContainersStart,
            "/containers/{name:.*}/stop":    postContainersStop,
            "/containers/{name:.*}/wait":    postContainersWait,
            "/containers/{name:.*}/resize":  postContainersResize,
            "/containers/{name:.*}/attach":  postContainersAttach,
            "/containers/{name:.*}/copy":    postContainersCopy,
            "/containers/{name:.*}/exec":    postContainerExecCreate,
            "/exec/{name:.*}/start":         postContainerExecStart,
            "/exec/{name:.*}/resize":        postContainerExecResize,
            "/containers/{name:.*}/rename":  postContainerRename,
    },
    "DELETE": {
            "/containers/{name:.*}": deleteContainers,
            "/images/{name:.*}":     deleteImages,
    },

