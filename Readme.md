# IM 

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)


此项目基于公司目前业务场景抽离的基础部分进行开源，此部分为用户与服务器维持基础连接，
im是一个通用的服务组件，且每个公司都希望拥有更好的用户体验。本项目会随着公司业务不断
的迭代，后续更新版本也会在github上更新


## 目录
 - [背景](##背景)
 - [安装](##安装)
 - [快速开始](##快速开始)
 - [性能](##性能)
    - [连接保持](##连接保持)
    - [消息下推](##消息下推)
    - [消息广播](##消息广播)
 
 
## 背景
  在最开始公司场景应用过程中只是做了很简单的单点im服务器，随着公司业务需求扩大，
  业务版本迭代加快，当时的im服务和业务耦合程度非常的深入，基本没有分离，每次更新
  就需要停服更新,后面构想将IM做为一个公共组件进行封装，业务层单独抽离，im服务就
  应运而生,这个版本是第一个版本的IM需求。
  
  在这个im项目上只需要专注于：
  - `用户链接保持`
  - `消息编解码`
  - `消息下推`
  - `在线状态维护`
  
  整个im服务都是不断的更新升级这几个基础工作。后续压力测试也将不断的围绕这几个
  功能进行测试
## 安装
  `go get github.com/mongofs/im`
  
   im 最多支持最新的两个go版本，请保证你本地的版本号
   
   
## 快速开始
你只需要在你的项目中使用下面这段代码，你就可以启动一个简易的IM聊天服务器，http默认端口是：8080，
RPC 服务端口默认是：8081，访问`ws://localhost:8080/chat?token=12345`即可建立连接。
```
    bs := im.New()
	err := bs.Run()
	if err !=nil {
		log.Fatal(err)
	}
	bs.Close()
```


## 关于选项
所有有关im的选项设置都放在可以使用With+ 等方式引出，由于每个公司业务场景不一样，比如有的公司心跳包只是
一个空包，有的公司的心跳包会携带当前的用户的数据同步id等等，基于这类业务不同在设计default结构体的时候
只能以最简单的方式。下面有两种选项建议设置了：建议开启就设置的选项。
#### 必选选项
- WithServerHTTPPort ：设置服务的HTTP端口
- WithServerRPCPort ：设置服务的RPC端口
- WithServerValidate ：注册函数，用户在与服务建立连接的时候，如何进行鉴权
- WithServerReceive ：注册函数，用户与服务器建立连接，当用户通过websocket发送消息，就会走到此处
- WithClientHeartBeatInterval ：设置心跳，设置用户与服务器之间的心跳间隔，可以在注册函数WithServerReceive上进行区分心跳包还是业务逻辑。

#### 可以进行性能调优项
- WithServerBucketNumber ：性能调优，由于对用户存储是使用的map，本身在并发访问的时候会出现并发问题，就将map进行分片锁的粒度减小提高性能
可以通过测试此参数进行观测具体压测结果，目前默认的参数是64。
- WithClientReaderBufferSize ： 用户的websocket预读的缓存长度。
- WithClientWriteBufferSize ： 用户的websocket写入的缓存长度。
- WithClientBufferSize ： 此参数是用户的预存buffer，在生产环境中用户会存在网络状态不好、网络拥塞等状态，导致数据写入过慢，尤其在广播业务
场景中我们需要将所有用户都进行通知一遍，防止某个用户的状态写入超时阻碍写入线程，就需要加一个buffer，建议长度在10以内。
- WithClientMessageType：设置用户的信息类型，就是文本类型还是数据流。
- WithClientProtocol：设置用户下发消息的具体协议：比如是json 还是pb ，在v1.01版本考虑针对每个用户的消息都可以设置。
- WithBucketSize：可以结合WithServerBucketNumber 来设置，比如你预计单机承载10000人的连接推送，那么10000 = bucketNumber * bucketSize
，由于底层是用hash表存放的，可以尽量设置大一点的bucketSize来避免扩容带来的性能损耗
- 











