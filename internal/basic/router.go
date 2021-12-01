package basic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"websocket/internal/basic/client"
	"websocket/internal/basic/recieve"
)

func (s *BasicServer) initRouter(middlewares ...gin.HandlerFunc)error{
	//分组创建路由
	s.http.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	s.http.Use(middlewares...)
	s.http.GET("/conn", s.ConnectSocket)
	return nil
}



// create the connection
func (s *BasicServer) ConnectSocket(ctx *gin.Context){

	token:= ctx.Query("token")
	if token == "" {
		return
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
		fmt.Println(err)
		return

	}
	// validate
	if err !=nil {
		cli.Send([]byte("用户非法"))
		cli.Offline()
		return
	}
	bs.Register(cli,token)
}