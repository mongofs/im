package im

import (
	"context"
	im "github.com/mongofs/api/im/v1"
	log "github.com/sirupsen/logrus"
	"time"
)


func (s *ImSrever) Ping(ctx context.Context, empty *im.Empty) (*im.Empty, error) {
	log.Infof(" |%c[1;40;32m RPC-%v   %c[0m",0x1B,"Ping",0x1B)
	return nil,nil
}


func (s *ImSrever) Onliens(ctx context.Context, empty *im.Empty) (*im.OnlinesReply, error) {
	log.Infof(" |%c[1;40;32m RPC-%v   %c[0m",0x1B,"Onliens",0x1B)
	num := s.ps.Load()
	req := & im.OnlinesReply{Number: num}
	return req,nil
}


func (s *ImSrever) SendMessage(ctx context.Context, req *im.SendMessageReq) (*im.SendMessageReply, error) {
	start  := time.Now()
	bs:= s.bucket(req.Token)
	err := bs.Send(req.Data,req.Token,false)
	escape := time.Since(start)
	log.Infof(" |%c[1;40;32m RPC-%v | %v  %c[0m",0x1B,"SendMessage",escape,0x1B)
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
	log.Infof(" |%c[1;40;32m RPC-%v | %v  %c[0m",0x1B,"SendMessageToMultiple",escape,0x1B)
	return &im.SendMessageReply{},err
}


func (s *ImSrever) Broadcast(ctx context.Context, req *im.BroadcastReq) (*im.BroadcastReply, error) {
	start  := time.Now()
	for _,v :=range s.bs{
		 v.BroadCast(req.Data,false)
	}
	escape := time.Since(start)
	log.Infof(" |%c[1;40;32m RPC-%v | %v  %c[0m",0x1B,"Broadcast",escape,0x1B)
	return &im.BroadcastReply{},nil
}


