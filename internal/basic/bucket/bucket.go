package bucket

import (
	"websocket/internal/basic/client"
)

type Bucketer interface {
	// send data to someone
	// ACK indicates whether the message is a receipt message
	Send(data []byte, token string ,Ack bool)error

	// Send messages to all online users
	SendAll(data []byte ,Ack bool)

	// Kick users offline
	OffLine(token string)

	// Register user to basket
	Register(cli client.Clienter,token string)error

	//Judge whether the user is online
	IsOnline(token string)bool
}