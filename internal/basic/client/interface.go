package client


type Clienter  interface {
	Send([]byte,...int64)
	Offline()
}