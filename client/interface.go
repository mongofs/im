package client


type Clienter  interface {
	// 调用此方法可以给当前用户发送消息
	Send([]byte,...int64)error
	// 调用此方法可以下线用户
	Offline()
	//
	ResetHeartBeatTime()

	LastHeartBeat()int64
}


