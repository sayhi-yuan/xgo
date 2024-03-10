# xgo
go的基础框架

### 1. 项目环境配置
#### 1.1 配置git
    1. 配置私有库
        go env -w GOPRIVATE=git.syreads.com
    2. 配置git账户
        git config user.username ""
        git config user.password ""
    3. 复制对应环境的配置文件，比如dev-config.yaml为config.yaml
    4. 在项目根目录下，执行构建命令 go build ./cmd/http/main.go ./bin/http

### 2. 项目要求
    1. golang 版本 1.20
    2. 文件夹、文件的命名，尽量一个单词说明，如果多个单词，则全小写，单词之间不加任何间隔符

### 3. 目录结构
    - bin               存放可执行文件的目录，可以不需要，根据自己的需求而定
    - bootstrap         项目配置环境的引导目录
    - cmd               不同场景的编译，目前支持web和定时任务
        - http          web main函数入口
        - job           定时任务main函数入口
    - config            配置目录
    - core              主要依赖存放地
    - deploy            项目文件，包含构建镜像、打包程序、ks8部署配置
    - internal          项目的业务代码模块
        - api           调用非本项目的http接口
        - cache         所有缓存
        - constant      项目常量
        - exception     异常
        - logic         公共的逻辑
        - middleware    中间件
        - models        模型，对应数据库的表。不能添加业务代码，但是可以添加修改器，访问器。模型关联
        - modules       业务模块
            - 模块名称
                - dto       和前端交互的参数，包含字段的验证
                - handel    接收参数，调用服务层，类似于controller
                    - 每个handle需要继承BaseHandle
                - service   服务层
                    - impl  接口实现层
                    - 服务模块
                        - init.go 服务的定义，以及服务公共方法
                        - 方法.go 建议一个方法一个文件
                        - 如果需要一些设计模式的引入，则继续加文件，但是不能被服务之外的地方引入
                    - 各个服务层的接口，实例化接口
                - router.go 模块路由，注入handle中的方法    
    - log               日志输出，支持控制台和日志文件
    - utils             工具方法的东西
    - config.yaml       配置文件
    - Jenkinsfile       ci/cd的流水线


### 4. 开发建议
    1. 函数或方法中如果有error，不能丢弃。比如有error的返回值，且放error必须在多返回值的最后一个
    2. 不同模块之间调用，尽量使用接口与实现分离的方式。
    3. 调用外部的api的代码，全部写在internal/api中，并处理好数据。业务代码只调用api成品
    4. 缓存相关的数据，全部写到cache中，业务只调用成品。
    5. models中，只能写和数据表相关的逻辑。不能添加业务逻辑
    6. handle请求到service之间的逻辑，例如，参数绑定、用户信息获取等
    7. 服务层代码采用接口与实现分离，接口在service下，实现在impl下。每个独立的方法，一个文件。struct和共有的私有方法，写在init.go下；
    8. 公共的逻辑放到logic中，比如多个service使用的方法，要提到logic中
    9. 需要认证的接口，都需要使用包装后的*core.Context。service的方法中，必须传递上下文
    10.参数的传递，尽量使用struct。在dto中定义好
    11.前端传递的参数，必须用struct接收。尽量不要使用map，除非struct干不了
    