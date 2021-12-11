package bucket

import (
	"github.com/mongofs/im/client"
	"time"
)

func (h *bucket) start (){
	go h.monitor()
	go h.keepAlive()
}


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
				//todo
				cancelClis := []client.Clienter{}
				now := time.Now().Unix()
				b.rw.Lock()
				for _, cli := range b.clis {
					if now-cli.LastHeartBeat() < 2*b.opts.HeartBeatInterval {
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



func (h *bucket)delUser(token string){
	h.rw.Lock()
	defer h.rw.Unlock()
	delete(h.clis,token)
	h.np.Add(-1)
	if h.opts.callback !=nil {
		h.opts.callback()
	}
}

