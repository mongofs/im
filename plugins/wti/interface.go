// WTI 全称为 websocket Target interface ，就是长连接标记接口，目前系统提供的解决方案为
package wti

import "github.com/mongofs/im/client"

//  WebSocket Target Interface
type WTI interface {
	// 给用户打上标签
	SetTAG(cli *client.Cli, tag ...string)

	// 如果用户下线将会通知调用这个方法
	Update(token ...string)

	// 广播到包含tag 对象
	BroadCast(content []byte, tag ...string)

	// 广播所有内容
	BroadCastByTarget(targetAndContent map[string][]byte)

	// 获取某个用户的所有的tag
	GetClienterTAGs(token string)[]string

	// 获取到TAG 的创建时间，系统会判断这个tag创建时间和当前人数来确认是否需要删除这个tag
	GetTAGCreateTime(tag string)int64

	// 获取到TAG 的在线人数，系统会判断这个tag如果没有在线人数为0 且创建时间大于MAX-wti-create-time ,这个tag就会被回收
	GetTAGClients (tag string)int64

	// 回收TAG ,im 主线程会根据GetTAGCreateTime 和 GetTAGClients 进行数据回收，回收也会调用此方法
	RecycleTAG (tag string)
}

// 其他地方将调用这个变量，如果自己公司实现tag需要注入在程序中进行注入
var Factory WTI = newwti()


func Inject (wti WTI) {
	Factory = wti
}


