# job-center
![Build Status](https://travis-ci.org/koko990/service-center.svg?branch=master)
## 概述
job-center设计初衷为任务调度、触发、管控的平台，job-center以go语言包方式的方式加载，
将job的入口方法传递给job-center编译成二进制程序执行，job-center在一个二进制程序内部代理job对外吐
的消息，并根据消息触发响应方法，消息按照严重等级划分为5个等级，对应触发的5种监控方法；在二进制程序的外部
job-center向etcd注册信息，并通过etcd完成job服务治理与调度。

## job的触发机制：
每个job都有计时器，皆可以固定时间间隔进行周期运行，job可由多种方式触发，如restapi发送开关信息、etcd，rpc等消息触发
可通过etcd改变job配置信息，设置周期参数。

## 特点
job-center可以通过go语言定义复杂的job，适合完成定期、多触发的任务模型，如CICD任务、爬虫程序的执行、大数据的定期计算、人工智能的算法训练等等，配合
容器云进行高并发处理事务。

## 局限性
由于golang语言为静态语言，因此必须要进行编译执行，当然业务写成动态脚本语言，如python、js等，通过本地文件搜索，或者etcd等系统
和job-center进行job日志的监控触发，可进行这一机制的探索，让job-center具备动态化

## 功能特性
### 1.  可通过api对job简单CRUD操作。
### 2。 通过rpc、rest、watch etcd等方式对job进行动态启停
### 3。 支持分布式功能，通过分布式锁保证job只会触发一次，通过算法控制分布式执行任务
### 4。 调度器监控cpu、memory等，对job进行弹性伸缩
### 5。 支持邮件报警、状态监控，实时查看任务数量、状态等。
### 6。 支持任务编排，控制任务执行的先后顺序
### 7。 支持分布式任务注册，将任务进行队列化存储至中心。
### 8。 分布式任务执行可进行丰富的路由策略

## 架构图
<img alt="job-center" src="docs/img/job_architecture.png">
