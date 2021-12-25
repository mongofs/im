package bucket

import (
	"github.com/mongofs/im/client"
	"go.uber.org/atomic"
	"time"
)

func (h *bucket) start (){
	go h.monitor()
	go h.keepAlive()
}

var temcounter  atomic.Int64

func (h *bucket)monitor (){
	if h.opts.ctx !=nil {
		for  {
			select {
			case token :=<- h.closeSig	:
				h.delUser(token)
			case <- h.opts.ctx.Done():
				return
			}
		}
	}
	for  {
		select {
		case token :=<- h.closeSig	:
			h.delUser(token)
		}
	}
}



func (b *bucket)keepAlive (){
	if b.opts.ctx !=nil {
		for {
			select {
			case <-b.opts.ctx.Done():
				return
			default:
				cancelClis := []client.Clienter{}
				now := time.Now().Unix()
				b.rw.Lock()
				for _, cli := range b.clis {
					// 如果心跳间隔 时间超过两个心跳包的时间，那么默认用户连接不可用
					if now-cli.LastHeartBeat() > 2*b.opts.HeartBeatInterval {
						continue
					}
					cancelClis = append(cancelClis,cli)
				}
				b.rw.Unlock()
				for _,cancel := range cancelClis{
					cancel.Offline()
				}
			}
			time.Sleep(10 * time.Second)
		}
	}

	for {
		now := time.Now().Unix()
		b.rw.Lock()
		for _, cli := range b.clis {
			if now-cli.LastHeartBeat() < 2*b.opts.HeartBeatInterval {
				continue
			}
			cli.Offline()
		}
		b.rw.Unlock()
		time.Sleep(10 * time.Second)
	}
}


func (h *bucket)delUser(token string) {
	h.rw.Lock()
	delete(h.clis, token)
	h.rw.Unlock()
	h.np.Add(-1)
	if h.opts.callback != nil {
		h.opts.callback()
	}
}


