package im

import (
	"context"
	"errors"
	"github.com/mongofs/im/client"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)


var ErrTokenIsNil =errors.New("basic : token can't be nil")

func (s *ImSrever) initRouter()error{
	//分组创建路由
	s.http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		res := &Response{
			w:      writer,
			Status: 403,
			Data:   "ok",
		}
		res.SendJson()
	})
	s.http.HandleFunc("/conn", s.Connection)
	return nil
}

// create  connection
func (s *ImSrever) Connection(writer http.ResponseWriter, request *http.Request){
	now :=time.Now()
	defer func() {
		escape := time.Since(now)
		log.Infof(" | %v|%c[1;40;32m HTTP-%v |%v  %c[0m| %v",
			request.RemoteAddr,0x1B,request.Method,escape,0x1B,request.URL)
	}()

	res := &Response{
		w:      writer,
		Data:   nil,
		Status: 200,
	}
	if request.ParseForm() != nil {
		res.Status = 400
		res.Data = "connection is bad "
		res.SendJson()
		return
	}

	token:= request.Form.Get("token")
	if token == "" {
		res.Status=400
		res.Data = "token validate error"
		res.SendJson()
	}
	// validate token
	bs:= s.bucket(token)
	ch := bs.NotifyBucketConnectionIsClosed()
	cli ,err := client.New(
		client.WithContext(context.Background()),
		client.WithReader(request),
		client.WithWriter(writer),
		client.WithUserToken(token),
		client.WithNotifyCloseChannel(ch),
		client.WithReceiveFunc(s.recevier.Handle),
		client.WithAgreement(s.agreement))
	if err !=nil {
		res.Status=400
		res.Data = err.Error()
		return
	}
	if err := s.validate.Validate(token);err !=nil {
		cli.Send([]byte("User token validate failed "))
		cli.Offline()
		return
	}
	if err := bs.Register(cli,token);err !=nil {
		cli.Send([]byte(err.Error()))
		cli.Offline()
	}
}



