package client

import (
	"context"
	"net/http"
)

type OptionFunc func(cli *client)

func WithWriter (writer http.ResponseWriter) OptionFunc {
	return func(cli *client) {
		cli.writer =writer
	}
}


func WithReader (r *http.Request) OptionFunc {
	return func(cli *client) {
		cli.reader =r
	}
}


func WithContext (ctx context.Context) OptionFunc {
	return func(cli *client) {
		cli.ctx =ctx
	}
}


func WithUserToken (token string) OptionFunc {
	return func(client *client) {
		client.token = token
	}
}



func WithReceiveFunc (f func(cli Clienter,data []byte)) OptionFunc {
	return func(client *client) {
		client.handleReceive =f
	}
}


func WithNotifyCloseChannel(ch  chan<- string) OptionFunc {
	return func(cli *client) {
		cli.closeSig = ch
	}
}