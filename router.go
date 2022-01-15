package im

import (
	"errors"
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
		s.opt.ServerLogger.Infof("im/router : %s create %s  cost %v  url is %v ", request.RemoteAddr,request.Method,escape,request.URL)
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
	cli,err := bs.CreateConn(writer,request,token,s.opt.ServerReceive)
	if err !=nil {
		res.Status=400
		res.Data = err.Error()
		return
	}
	// validate failed
	if err := s.opt.ServerValidate.Validate(token);err !=nil {
		s.opt.ServerValidate.ValidateFailed(err,cli)
		return
	}else {
		//validate success
		s.opt.ServerValidate.ValidateSuccess(cli)
	}


	// register to data
	if err := bs.Register(cli,token);err !=nil {
		cli.Send([]byte(err.Error()))
		cli.Offline()
	}
}



