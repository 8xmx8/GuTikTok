<p align="center">
  <a href="https://github.com/Ocyss/Douyin">
    <img src="https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Q/douyin/logo.svg" alt="Logo" width="180" height="180">
  </a>

  <h1 align="center">极简版抖音</h1>
<p align="center"> 开发环境1024code,字节官方提供1G的代码空间
 这是一个代码空间的副本，由于开发环境原因，特将项目代码拷贝下来，留作纪念</p>

  <p align="center">
    一个字节青训营的实战项目
	<br/>
	开始于2023.7.24 结束于2023.8.20
    <br/>
     <br/>
    <a href="https://github.com/8xmx8/GuTikTok/issues">报告Bug</a>
    <a href="https://github.com/8xmx8/GuTikTok/issues">提出新特性</a>
</p>

## 技术栈

#### 后端 Golang 1.20

- Gin [(Web 框架)](https://gin-gonic.com/zh-cn/)
- Grpc[(RPC 框架)]()
- Proto3[(IDL语言)]()
- Consul[(服务发现)]()
- GORM [(ORM)](https://gorm.io/zh_CN/)
- MySQL [(数据库)]()
- Redis [(缓存)]()
- Pyroscope[(性能分析)]()
- Elasticsearch[(搜索引擎)]()
- RabbitMQ[(消息中间件)]()
- OpenTelemetry,VictoriaMetrics[(可观测性)]()
- Ffmpeg[(视频处理)]()
- Gorse[(推荐系统)]()

## 部署方法

#### 1. 确保以下服务为开启状态

- Consul
- Jaeger
- Mysql
- Redis
- Pyroscope
- Prometheus

推荐使用docker进行部署：
>  Jaeger 和 Pyroscope 以及 Prometheus 推荐使用 Docker 进行部署，部署命令如下：
> 
>  docker run -d --name pyroscope -p 4040:4040 pyroscope/pyroscope:latest server
>  
>  docker run -d --name=jaeger -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest
> 
>  docker run --name prometheus -d -p 127.0.0.1:9090:9090 prom/prometheus

#### 2.项目结构

```
├─src # GuTikTok 源代码
│  ├─constant # 常量，用于定义服务预设信息等
│  │  ├─config # 为项目提供配置读取说明
│  │  └─strings # 提供编码后的信息，用于定义常量信息
│  ├─idl # idl 说明文件
│  ├─models # 用于存储通用数据模型
│  ├─rpc # gRPC 生成文件
│  ├─services # 服务，下为横向扩展的服务
│  │  ├─auth # Auth 鉴权 / 登录服务
│  │  └─health # *唯一一个非横向扩展服务，用于注册到其他服务中，提供 consul 健康检查的功能
│  ├─storage # 存储模块，暂时缺少 RabbitMQ 对接模块，需要由视频相关业务开发组制作
│  │  ├─database # 数据库模块，对接 mysql
│  │  ├─file # 二进制存储模块，目前只有 fs 模块
│  │  └─redis # Redis 模块，对接 Redis
│  ├─utils # 通用问题
│  │  ├─consul # Consul 服务，用于向 Consul 注册服务
│  │  ├─interceptor # 拦截器，用于切片某一个方法或过程
│  │  ├─logging # 日志
│  │  └─trace # 链路追踪
│  └─web # 网页服务
│      ├─about # *About 服务，非正式业务，仅供测试
│      ├─auth # Auth 服务，提供 /douyin/user * 接口
│      ├─authmw # Auth 鉴权中间件，非服务
│      ├─middleware # Middle Ware 中间件，为除了 Auth MW 以外的中间件服务
│      └─models # 网站模型
└─test # 单元测试
├─rpc # GRPC 单元测试
└─web # 网页单元测试
```

#### 3.项目启动
> web端口:37000

在启动web服务之前，请把所有的微服务开启
