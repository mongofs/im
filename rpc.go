package im

import (
	"context"
	im "github.com/mongofs/api/im/v1"
)


func (s *ImSrever) Ping(ctx context.Context, empty *im.Empty) (*im.Empty, error) {
	return nil,nil
}


func (s *ImSrever) Onliens(ctx context.Context, empty *im.Empty) (*im.OnlinesReply, error) {
	num := s.ps.Load()
	req := & im.OnlinesReply{Number: num}
	return req,nil
}


func (s *ImSrever) SendMessage(ctx context.Context, req *im.SendMessageReq) (*im.SendMessageReply, error) {
	bs:= s.bucket(req.Token)
	return &im.SendMessageReply{},bs.Send(req.Data,req.Token,false)
}


func (s *ImSrever) Broadcast(ctx context.Context, req *im.BroadcastReq) (*im.BroadcastReply, error) {
	for _,v :=range s.bs{
		go v.BroadCast(req.Data,false)
	}
	return &im.BroadcastReply{},nil
}