package rooms

import (
	"github.com/mongofs/im/client"
	"go.uber.org/atomic"
	"sync"
	"time"
)

type roomSet struct {
	createM func(RooID string) room
	rw      *sync.RWMutex
	rooms   map[string]room
	clear   time.Duration

	inRoom *atomic.Int64
}



func (r *roomSet) GetRoomUserNumber() int32 {
	panic("implement me")
}


func (r *roomSet) monitor() {
	for {
		r.rw.Lock()
		r.inRoom.Store(0)
		for k, v := range r.rooms {
			// delete the room
			counter := v.GetRoomUserNumber()
			if counter == 0 && time.Now().Unix()-v.GetRoomCreateTime() > 60*60*2 {
				delete(r.rooms, k)
				continue
			}
			r.inRoom.Add(int64(counter))
		}
		r.rw.Unlock()


		time.Sleep(r.clear * time.Second)
	}
}



func (r *roomSet) AddClientToRoom(token string, conn client.Clienter, RoomID string) {
	r.rw.Lock()
	tem, ok := r.rooms[RoomID]
	r.rw.Unlock()
	if !ok {
		temRoom := r.createM(RoomID)
		temRoom.AddUser(token, conn)
		r.rw.Lock()
		r.rooms[RoomID] = temRoom
		r.rw.Unlock()
	}
	tem.AddUser(token, conn)
}



func (r *roomSet)GetRoomOnlineList(){



}




func (r roomSet) DelClientFromRoom(token string, RoomID string) {
	r.rw.Lock()
	tem, ok := r.rooms[RoomID]
	r.rw.Unlock()
	if !ok {
		return
	}
	tem.DelUser(token)
}



func (r roomSet) PushDataToRoom(data []byte, RoomID string, token ...string) {
	r.rw.Lock()
	tem, ok := r.rooms[RoomID]
	r.rw.Unlock()
	if !ok {
		return
	}

	if len(token) == 0 {
		tem.PushData(data) //全房间推送
	} else {
		tem.PushDataToPointedUser(data, token...) // 推送给指定用户
	}
}
