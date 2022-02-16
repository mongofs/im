package client

import "net/http"

type Clienter interface {

	// 调用此方法可以给当前用户发送消息
	Send([]byte, ...int64) error

	// 用户下线
	Offline()

	// 重置用户的心跳
	ResetHeartBeatTime()

	// 获取用户的最后一次心跳
	LastHeartBeat() int64

	// 获取用户的token
	Token() string

	// 获取到用户的请求的链接
	Request()*http.Request
}
