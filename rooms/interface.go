package rooms

import "github.com/mongofs/im/client"

type room interface {
	AddUser(token string, clienter client.Clienter) // 将用户添加到某个房间

	DelUser(token string) // 将用户从房间删除

	PushData(data []byte) int // 推送消息给房间,并返回推送在线用户数量

	PushDataToPointedUser(data []byte, token ...string) []string // 将消息推送给房间指定用户

	GetRoomUserNumber() int32 //获取当前房间人数：注意特指当前服务的机器，不能代表房间所有用户，建议通过redis incr 存储

	GetRoomCreateTime() int64 // 获取当前房间创建时间

}

// 房间管理器
type RoomSet interface {

	// add user to room
	AddClientToRoom(token string, conn client.Clienter, RoomID string)
	// del user to room
	DelClientFromRoom(token string, RoomID string)
	// push data to room
	PushDataToRoom(data []byte, RoomID string, token ...string)
	// get room usernumber
	GetRoomUserNumber() int32
}
