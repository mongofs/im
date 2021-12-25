package client


type Clienter  interface {
	// 调用此方法可以给当前用户发送消息
	Send([]byte,...int64)error
	// 用户下线
	Offline()
	//
	ResetHeartBeatTime()

	LastHeartBeat()int64

	Token()string
}


