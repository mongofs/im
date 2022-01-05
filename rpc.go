package im

import (
	"context"
	"errors"
	"fmt"
	im "github.com/mongofs/api/im/v1"
	"time"
)


func (s *ImSrever) Ping(ctx context.Context, empty *im.Empty) (*im.Empty, error) {
	s.opt.ServerLogger.Infof(" im/rpc : called  %v method  ","Ping")
	return nil,nil
}


func (s *ImSrever) Onliens(ctx context.Context, empty *im.Empty) (*im.OnlinesReply, error) {
	s.opt.ServerLogger.Infof(" im/rpc : called  %v method  ","Onliens")
	num := s.ps.Load()
	req := & im.OnlinesReply{Number: num}
	return req,nil
}


func (s *ImSrever) SendMessage(ctx context.Context, req *im.SendMessageReq) (*im.SendMessageReply, error) {
	start  := time.Now()
	bs:= s.bucket(req.Token)
	err := bs.Send(req.Data,req.Token,false)
	escape := time.Since(start)
	s.opt.ServerLogger.Infof(" im/rpc : called  %v method cost time %v ","SendMessage",escape)

	return &im.SendMessageReply{},err
}


// 相同消息发送给多个用户
func (s *ImSrever) SendMessageToMultiple(ctx context.Context, req *im.SendMessageToMultipleReq) (*im.SendMessageReply, error) {
	start  := time.Now()
	var err error
	for _,token := range req.Token{
		bs:= s.bucket(token)
		err = bs.Send(req.Data,token,false)
	}
	escape := time.Since(start)
	s.opt.ServerLogger.Infof(" im/rpc : called  %v method cost time %v ","SendMessageToMultiple",escape)
	return &im.SendMessageReply{},err
}

// 广播消息给用户
func (s *ImSrever) Broadcast(ctx context.Context, req *im.BroadcastReq) (*im.BroadcastReply, error) {
	if len(s.buffer) *10  > 8* cap(s.buffer){
		return nil,errors.New(fmt.Sprintf("im/rpc: too much message ,buffer length is  %v but cap is %v",len(s.buffer),cap(s.buffer)))
	}
	start  := time.Now()
	s.buffer <- req
	escape := time.Since(start)
	s.opt.ServerLogger.Infof(" im/rpc : called  %v method cost time %v ","Broadcast",escape)
	return &im.BroadcastReply{
		Size: int64(len(s.buffer)),
	},nil
}


