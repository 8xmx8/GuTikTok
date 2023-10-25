# GuTikTok
<p>字节青训营项目重构<a href="https://github.com/Godvictory/douyin">
&nbsp;第一版地址 </a> &nbsp;&nbsp; <a href="https://github.com/Godvictory/douyin#readme">第一版文档</a></p>
<br/>
<p align="center">
  <a href="https://github.com/Ocyss/Douyin">
    <img src="https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Q/douyin/logo.svg" alt="Logo" width="180" height="180">
  </a>

  <h1 align="center">极简版抖音</h1>
  <p align="center">
    一个字节青训营的实战项目
	<br/>
	开始于2023.7.24 结束于2023.8.20
    <br/>
     <br/>
    <a href="https://github.com/Godvictory/GuTikTok/issues">报告Bug</a>
    <a href="https://github.com/Godvictory/GuTikTok/issues">提出新特性</a>
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
- OpenTelemetry,VictoriaMetrics[(可观测性)]()
- Ffmpeg[(视频处理)]()
- Gorse[(推荐系统)]()
#### 前端 Vue.js 3

- Vite [(构建工具)](https://cn.vitejs.dev/)
- element-plus [(UI 库)](https://element-plus.org/zh-CN/)
- xgplayer [(西瓜播放器)](https://v2.h5player.bytedance.com/gettingStarted/)
- md-editor-v3 [(Markdown 编辑器)](https://www.wangeditor.com/)

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

#### 2. clone 项目
> git clone git@github.com:Godvictory/GuTikTok.git
#### 3.项目启动
> web端口:37000

在启动web服务之前，请把所有的微服务开启
