package bucket

import (
	"context"
	"errors"
	"github.com/mongofs/im/ack"
	"github.com/mongofs/im/client"
	"go.uber.org/atomic"
	"sync"
)


var (
	ErrUserExist =errors.New("hash : Cannot login repeatedly")
	ErrCliISNil  =errors.New("hash : cli is nil")
)

type hash struct {
	rw sync.RWMutex

	// Number of people
	np *atomic.Int64

	// users set
	users map[string]client.Clienter

	// size
	size int32

	// User offline notification
	closeSig chan string

	// User offline callback method
	offline func()

	// Ack map
	ack ack.Acker

	// context
	ctx context.Context
}

const DefaultBufferSize = 100


func New(opt ...OptionFunc) Bucketer {

	res := & hash{
		rw:       sync.RWMutex{},
		np:       &atomic.Int64{},
		closeSig: make(chan string,1),
	}

	if len(opt)>0 {
		for _,o := range opt{
			o(res)
		}
	}
	if res.size ==0 {
		res.size = DefaultBufferSize
	}
	res.users = make(map[string]client.Clienter,res.size)

	res.start()
	return res
}


func (h *hash)randId()int64{
	return 0
}


func (h *hash) Onlines()int64 {
	return h.np.Load()
}


func (h *hash)Flush(){
	h.rw.RLock()
	defer h.rw.RUnlock()
	h.np.Store(int64(len(h.users)))
}



func (h *hash) send (cli client.Clienter,token string,data []byte,ack bool)error{
	if ack {
		sid := h.randId()
		if err := h.ack.AddMessage(token,sid,data);err !=nil{
			return err
		}
		cli.Send(data,sid)
	}else{
		cli.Send(data)
	}
	return nil
}



func (h *hash) Send(data []byte, token string, Ack bool) error{
	h.rw.RLock()
	defer h.rw.RUnlock()
	if cli ,ok:= h.users[token];!ok{
		return ErrCliISNil
	}else {
		return h.send(cli,token,data,Ack)
	}
}



func (h *hash) BroadCast(data []byte, Ack bool) {
	h.rw.RLock()
	defer h.rw.RUnlock()
	for token,cli := range h.users{
		h.send(cli,token,data,Ack)
	}
}



func (h *hash) OffLine(token string) {
	h.rw.RLock()
	defer h.rw.RUnlock()
	cli := h.users[token]
	cli.Offline()
}




func (h *hash) Register(cli client.Clienter,token string) error {
	if cli == nil  {
		return ErrCliISNil
	}
	h.rw.Lock()
	defer h.rw.Unlock()
	if _,ok := h.users[token]; ok {
		return ErrUserExist
	}
	h.users[token] = cli
	h.np.Add(1)
	return nil
}




func (h *hash) IsOnline(token string) bool {
	h.rw.RLock()
	defer h.rw.RUnlock()
	if _,ok:= h.users[token];ok {
		return true
	}
	return false
}



func (h *hash)NotifyBucketConnectionIsClosed()chan <- string{
	return h.closeSig
}
