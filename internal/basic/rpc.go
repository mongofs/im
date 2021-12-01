package basic

import (
	"context"

	im "github.com/mongofs/api/im/v1"
)


func (s *BasicServer) Ping(ctx context.Context, empty *im.Empty) (*im.Empty, error) {
	return nil,nil
}


func (s *BasicServer) Onliens(ctx context.Context, empty *im.Empty) (*im.OnlinesReply, error) {
	req := & im.OnlinesReply{Number: s.ps.Load()}
	return req,nil
}


func (s *BasicServer) SendMessage(ctx context.Context, req *im.SendMessageReq) (*im.SendMessageReply, error) {
	bs:= s.bucket(req.Token)
	return nil,bs.Send(req.Data,req.Token,false)
}


func (s *BasicServer) Broadcast(ctx context.Context, req *im.BroadcastReq) (*im.BroadcastReply, error) {
	for _,v :=range s.bs{
		v.BroadCast(req.Data,false)
	}
	return nil,nil
}