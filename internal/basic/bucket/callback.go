package bucket

import (
	"fmt"
)


func (h *hash) start (){
	go h.monitor()
}


func (h *hash)monitor (){
	defer func() {
		if err := recover();err !=nil {
			fmt.Println( err )
		}
	}()

	if h.ctx !=nil {
		for  {
			select {
			case token :=<- h.closeSig	:
				h.delUser(token)
			case <- h.ctx.Done():
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



func (h *hash)delUser(token string){
	h.rw.Lock()
	defer h.rw.Unlock()
	delete(h.users,token)
	h.np.Add(-1)
	if h.offline !=nil {
		h.offline()
	}
}

