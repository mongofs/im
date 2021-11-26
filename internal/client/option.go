package client

import (
	"context"
	"net/http"
)

type Option interface {
	apply(cli *client)
}

type optionFunc func(cli *client)


func (o optionFunc)apply(cli *client){
	o(cli)
}



func ClientWithWriter (writer http.ResponseWriter) optionFunc {
	return func(client *client) {
		client.writer = writer
	}
}


func ClientWithReader (r *http.Request) optionFunc {
	return func(client *client) {
		client.reader = r
	}
}


func ClientWithContext (ctx context.Context) optionFunc {
	return func(client *client) {
		client.ctx = ctx
	}
}


func ClientWithUserToken (token string) optionFunc {
	return func(client *client) {
		client.token = token
	}
}
