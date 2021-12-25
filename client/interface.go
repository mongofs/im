package client


type Clienter  interface {
	// 调用此方法可以给当前用户发送消息
	Send([]byte,...int64)error
	// 调用此方法可以下线用户 ,默认请不用传入参数，不然可能导致panic
	Offline(forRetry ...bool)
	//
	ResetHeartBeatTime()

	LastHeartBeat()int64

	Token()string
}


