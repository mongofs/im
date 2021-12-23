package rooms

import (
	"github.com/mongofs/im/client"
	"go.uber.org/atomic"
	"sync"
	"time"
)

type MapRoom struct {
	rw  *sync.RWMutex
	ro  map[string]client.Clienter
	cot *atomic.Int32 //当前在线用户
	createTime int64
}



func (m *MapRoom) GetRoomCreateTime() int64 {
	return m.createTime
}

func NewMapRoom()room{
	return &MapRoom{
		rw:  &sync.RWMutex{},
		ro:  make(map[string]client.Clienter,20),
		cot: &atomic.Int32{},
		createTime: time.Now().Unix(),
	}
}


func (m *MapRoom) AddUser(token string, clienter client.Clienter) {
	m.rw.Lock()
	m.ro[token]= clienter
	m.rw.Unlock()

	m.cot.Inc()
}



func (m *MapRoom) DelUser(token string) {
	m.rw.Lock()
	delete(m.ro,token)
	m.rw.Unlock()
	m.cot.Dec()
}



func (m *MapRoom) PushData(data []byte) int{
	m.rw.RLock()
	for _,v:= range m.ro{
		v.Send(data) // todo 需要优化
	}
	m.rw.RUnlock()
}



func (m *MapRoom) PushDataToPointedUser(data []byte, token ...string)[]string {
	m.rw.RLock()
	for _,v:= range token{
		if temCli,ok:= m.ro[v];ok{
			temCli.Send(data)
		}
	}
	m.rw.Unlock()
}



func (m *MapRoom) GetRoomUserNumber() int32 {
	return m.cot.Load()
}



