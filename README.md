# Toy-Framework

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

短期内应该是不会有的

# 内容

## 数据结构

### Toy-Kv

toy-kv是一个并发安全的，千万吞吐的kv存储数据结构，采用了分段分表锁进行并发控制。

**博客文档**

+ [(分段分表锁)使用Golang实现一个并发安全的Map](https://www.sunist.cn/post/KeyValueStore-GolangImplement-3)

**视频教学**

1. [[Go] 从零开始实现一个每秒千万级操作的Key-Value存储 - Introduction](https://www.bilibili.com/video/BV1JG4y1e7NC)
2. [[Go] 从零开始实现一个每秒千万级操作的Key-Value存储 - CommonMap](https://www.bilibili.com/video/BV1BB4y167bY)
3. [[Go] 从零开始实现一个每秒千万级操作的Key-Value存储 - SimpleMap](https://www.bilibili.com/video/BV1bv4y1c7sg)
4. [[Go] 从零开始实现一个每秒千万级操作的Key-Value存储 - RWMap](https://www.bilibili.com/video/BV18d4y1K7rf)

# Contact

虽然应该没人会找我的，但还是留一个邮箱吧: [sunist's mail](mailto:sunist@mail.swu-acm.cn)