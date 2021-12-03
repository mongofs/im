package bucket

import (
	"websocket/client"
)

type Bucketer interface {
	// send data to someone
	// ACK indicates whether the message is a receipt message
	Send(data []byte, token string ,Ack bool)error

	// Send messages to all online users
	BroadCast(data []byte ,Ack bool)

	// Kick users offline
	OffLine(token string)

	// Register user to basket
	Register(cli client.Clienter,token string)error

	//Judge whether the user is online
	IsOnline(token string)bool


	Onlines()int64


	Flush()


	NotifyBucketConnectionIsClosed()chan <- string
}