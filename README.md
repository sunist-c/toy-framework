# Toy-Framework

[![Build Status](https://ci.swu-acm.cn/api/badges/sunist-c/toy-framework/status.svg?ref=refs/heads/main)](https://ci.swu-acm.cn/sunist-c/toy-framework)

toy-framework是[sunist-c](https://www.sunist.cn)在学习和开发的过程中，萌发的项目，主要的目的是从零开始，逐渐搭建一个HTTP/RPC服务框架。

在早期，本仓库可能一直起灰，因为sunist-c可能在极度愤怒的情况下大删大改，revert很多次，所以sunist-c会在toy-framework快完成的时候，再将其提交到GitHub

## 写在前面

本项目适用于新人从头开始使用go进行web后端开发

如果你需要一个能够生产的框架，请转到[go-gin-api](https://github.com/xinliangnote/go-gin-api)

如果你需要一些gin的开发样例，请转到[go-gin-example](https://github.com/eddycjy/go-gin-example)

## 说明

本项目的文档会在[sunist-c's blog](https://www.sunist.cn)进行更新，本仓库的readme也会贴出相应的链接

部分内容会在[SunistC - bilibili](https://b23.tv/HDAiCqq)进行视频教学，本仓库的readme也会贴出相应的链接

## Change Log

暂时应该不会有

# 内容

## 数据结构

### Toy-Kv (✅ Release)

toy-kv是一个并发安全的，千万吞吐的kv存储数据结构，采用了分段分表锁进行并发控制

> 这一节主要使用了互斥锁、读写锁以及分段式读写锁进行并发控制，适合未接触过并发编程的小伙伴入门

**博客文档**

+ [(分段分表锁)使用Golang实现一个并发安全的Map](https://www.sunist.cn/post/KeyValueStore-GolangImplement-3)

**视频教学**

1. [[Go] 从零开始实现一个每秒千万级操作的Key-Value存储 - Introduction](https://www.bilibili.com/video/BV1JG4y1e7NC)
2. [[Go] 从零开始实现一个每秒千万级操作的Key-Value存储 - CommonMap](https://www.bilibili.com/video/BV1BB4y167bY)
3. [[Go] 从零开始实现一个每秒千万级操作的Key-Value存储 - SimpleMap](https://www.bilibili.com/video/BV1bv4y1c7sg)
4. [[Go] 从零开始实现一个每秒千万级操作的Key-Value存储 - RWMap](https://www.bilibili.com/video/BV18d4y1K7rf)

### Toy-Rds (☑️ Developing)

toy-rds是一个使用go实现的类redis服务，使用epoll作为底层网络实现，可以做到百万级TPS

> 这一节主要使用了chanel、wait-group来进行并发控制，使用的是共享内存方式在不同go程之间传输数据，同时使用TCP向外部系统暴露接口，使用epoll进行网络并发控制

### Toy-Chan (⭕️ Planning)

toy-chan是一个并发安全的chanel，在原生chan的基础上增加了自动容量管理机制，支持自动扩容与自动释放多余容量

## 框架

### Toy-Http (☑️ Developing)

toy-http是一个基于gin的HTTP服务框架，在gin的基础上增加了依赖管理与参数管理，支持自动校验参数合法性，自动注入需要的依赖，自动生成状态与响应、编译时依赖检查、运行时动态依赖修改等功能

**博客文档**

+ 暂未更新博客

**视频教学**

1. [[Go] 从零开始实现一个HTTP框架 - Introduction](https://www.bilibili.com/video/BV1ra4y1f7ba)

### Toy-Rpc (❌ mkdir...)

toy-rpc是一个基于json-rpc的远端调用服务框架

### Toy-Mic (❌ mkdir...)

toy-mic是一个还没想好要怎么做的微服务框架

## 应用

### Toy-Crud (❌ mkdir...)

toy-crud是一个基于toy-http开发的基础CRUD应用，在给定Model的情况下可以快速生成CRUD接口与服务，除了定义Model以外无需任何额外代码

### Ceobebot-Backend-NG (⭕️ Planning)

[CeobeBot-Backend](https://github.com/ceobebot)是对话机器人CeobeBot的HTTP后端，为使用[Adachi-Bot](https://github.com/Arondight/Adachi-BOT)的CeobeBot前端提供基于HTTP的数据处理服务，目前正在使用toy-http进行重构

# Contact

虽然应该没人会找我的，但还是留一个邮箱吧: [sunist's mail](mailto:sunist@mail.swu-acm.cn)

# Thanks for supporting

1. [Southwest University ACM-Laboratory](https://github.com/swu-acm-lab): 提供了私有git仓库与ci/cd流水线
2. Southwest University Animal Science Research Institute: 基于toy-http前身开发后端服务，支持toy-http框架进行迭代开发
3. [Jecosine](https://github.com/jecosine)等: 为toy-framework提供测试与改进意见

# License

[MIT-LICENSE](./LICENSE)，你可以随便使用，但是还是希望你在使用本项目的相关内容的时候，能够带上项目链接或作者信息