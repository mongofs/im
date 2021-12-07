package im

import (
	"github.com/mongofs/im/recieve"
	 "github.com/mongofs/im/validate"
)

type Option func (b *ImSrever)


func WithHttpPort (httpPort string) Option {
	return func(b *ImSrever) {
		b.httpPort=httpPort
	}
}



func WithRpcPort (Port string) Option {
	return func(b *ImSrever) {
		b.rpcPort=Port
	}
}


func WithUsersValidater (validate validate.Validater) Option {
	return func(b *ImSrever) {
		b.validate= validate
	}
}



func WithClientReceiver  (receiver recieve.Receiver) Option {
	return func(b *ImSrever) {
		b.recevier = receiver
	}
}


func WithClientAgreement  (agreement int) Option {
	return func(b *ImSrever) {
		b.agreement =agreement
	}
}

func WithClientMessageSendMethod  (messageSendMethod int) Option {
	return func(b *ImSrever) {
		b.MessageSendMethod =messageSendMethod
	}
}