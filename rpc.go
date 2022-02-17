package im

import (
	"context"
	"errors"
	"fmt"
	im "github.com/mongofs/api/im/v1"
	"github.com/mongofs/im/plugins/wti"
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

// 在开发过程中存在IM需要版本共存的需求，比如我的协议替换了，但是如果im应用在App上面如何进行切换，这就是协议定制不合理的地方，但也需要
// IM 服务器在这个过程中做配合。
// IM 存在给用户分组的需求，所以我们在进行Broadcast 就必须进行用户的状态区分，所以前台需要对内容进行分组，传入的内容也需要对应分组
// 比如 v1 => string ，v2 => []byte，那么v1，v2 就是不相同的两个版本内容。在client上面可以设置用户的连接版本Version，建议在
// 使用用户

func (s *ImSrever) BroadcastByWTI(ctx context.Context, req *im.BroadcastByWTIReq) (*im.BroadcastReply, error) {
	var err error
	start  := time.Now()
	err = wti.BroadCastByTarget(req.Data)
	escape := time.Since(start)
	s.opt.ServerLogger.Infof(" im/rpc : called  %v method cost time %v ","BroadcastByWTI",escape)
	return &im.BroadcastReply{
		Size: int64(len(s.buffer)),
	},err
}

