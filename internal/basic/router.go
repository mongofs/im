package basic

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"websocket/internal/basic/client"
	"websocket/internal/basic/recieve"
)


var ErrTokenIsNil =errors.New("basic : token can't be nil")

func (s *BasicServer) initRouter(middlewares ...gin.HandlerFunc)error{
	//分组创建路由
	s.http.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	s.http.Use(middlewares...)
	s.http.GET("/conn", func(c *gin.Context) {
		if err := s.Connection(c); err != nil {
			fmt.Println(err)
			return
		}
	})
	return nil
}



// create the connection
func (s *BasicServer) Connection(ctx *gin.Context)error{

	token:= ctx.Query("token")
	if token == "" {
		return ErrTokenIsNil
	}
	// validate token
	bs:= s.bucket(token)
	ch := bs.NotifyBucketConnectionIsClosed()
	cli ,err := client.New(
		client.WithContext(ctx),
		client.WithReader(ctx.Request),
		client.WithWriter(ctx.Writer),
		client.WithUserToken(token),
		client.WithNotifyCloseChannel(ch),
		client.WithReceiveFunc(recieve.Handle))

	if err !=nil {
		return err

	}
	// ! validate
	if token =="error" {
		cli.Send([]byte("not ok "))
		cli.Offline()
		return err
	}

	return bs.Register(cli,token)
}