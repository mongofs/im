package bucket

import (
	"errors"
	"fmt"
	"github.com/mongofs/im/ack"
	"github.com/mongofs/im/client"
	log "github.com/sirupsen/logrus"
	"go.uber.org/atomic"
	"net/http"
	"sync"
)


var (
	ErrUserExist =errors.New("hash : Cannot login repeatedly")
	ErrCliISNil  =errors.New("hash : cli is nil")
)

type bucket struct {
	rw sync.RWMutex

	// Number of people
	np *atomic.Int64

	// users set
	clis map[string]client.Clienter

	// User offline notification
	closeSig chan string

	// Ack map
	ack ack.Acker


	opts * Option
}

func New(option *Option) Bucketer {
	res := & bucket{
		rw:       sync.RWMutex{},
		np:       &atomic.Int64{},
		closeSig: make(chan string,0),
		opts: option,
	}
	res.clis = make(map[string]client.Clienter,res.opts.BucketSize)
	res.start()
	return res
}


func (h *bucket)Flush(){
	h.rw.RLock()
	defer h.rw.RUnlock()
	h.np.Store(int64(len(h.clis)))
}


func(h *bucket)CreateConn(w http.ResponseWriter,r * http.Request,token string,handler client.Receiver)(client.Clienter,error){
	return  client.CreateConn(w , r ,
				h.closeSig,
				h.opts.ClientBufferSize,
				h.opts.MessageType,
				h.opts.Protocol,
				h.opts.ReaderBufferSize,
				h.opts.WriteBufferSize,
				token ,
				h.opts.ctx,
				handler)
}

func (h *bucket)randId()int64{
	return 0
}

func (h *bucket) Onlines()int64 {
	return h.np.Load()
}



func (h *bucket) send (cli client.Clienter,token string,data []byte,ack bool)error{
	if ack {
		sid := h.randId()
		if err := h.ack.AddMessage(token,sid,data);err !=nil{
			return err
		}
		return  cli.Send(data,sid)
	}else{
		return cli.Send(data)
	}
	return nil
}

func (h *bucket) Send(data []byte, token string, Ack bool) error{
	h.rw.RLock()
	cli ,ok:= h.clis[token];
	h.rw.RUnlock()
	if !ok{
		return ErrCliISNil
	}else {
		return h.send(cli,token,data,Ack)
	}
}

func (h *bucket) BroadCast(data []byte, Ack bool) error{
	counter := 0
	h.rw.RLock()
	for token,cli := range h.clis{
		err := h.send(cli,token,data,Ack)
		if err !=nil {
			counter ++
		}
	}
	h.rw.RUnlock()
	if counter !=0 {return fmt.Errorf("im/client : some user  can't arrive , the count is %v",counter)}
	return nil
}

func (h *bucket) OffLine(token string) {
	h.rw.RLock()
	defer h.rw.RUnlock()
	cli := h.clis[token]
	cli.Offline()
}

func (h *bucket) Register(cli client.Clienter,token string) error {
	if cli == nil  {
		return ErrCliISNil
	}
	h.rw.Lock()
	defer h.rw.Unlock()
	old,ok := h.clis[token];
	if ok {
		log.Infof(fmt.Sprintf("im: User token %s is online, but is trying to connect again",token))
		clienter ,_:= old.(*client.Cli)
		clienter.OfflineForRetry(true)
	}
	h.clis[token] = cli
	h.np.Add(1)
	return nil
}

func (h *bucket) IsOnline(token string) bool {
	h.rw.RLock()
	defer h.rw.RUnlock()
	if _,ok:= h.clis[token];ok {
		return true
	}
	return false
}


func (h *bucket)NotifyBucketConnectionIsClosed()chan <- string{
	return h.closeSig
}




