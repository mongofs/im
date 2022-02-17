// WTI 全称为 websocket Target interface ，就是长连接标记接口，目前系统提供的解决方案为
package wti

import (
	"errors"
	"github.com/mongofs/im/client"
	"go.uber.org/atomic"
)

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
	GetClienterTAGs(token string) []string

	// 获取到TAG 的创建时间，系统会判断这个tag创建时间和当前人数来确认是否需要删除这个tag
	GetTAGCreateTime(tag string) int64

	// 获取到TAG 的在线人数，系统会判断这个tag如果没有在线人数为0 且创建时间大于MAX-wti-create-time ,这个tag就会被回收
	GetTAGClients(tag string) int64

	// 调用方法的回收房间的策略
	FlushWTI()
}

// 其他地方将调用这个变量，如果自己公司实现tag需要注入在程序中进行注入
var factory WTI = newwti()
var isSupportWTI = atomic.NewBool(false)

func Inject(wti WTI) {
	factory = wti
}

func SetSupport (){
	isSupportWTI.Store(true)
}

var (
	ERRNotSupportWTI = errors.New("im/plugins/wti: not set the wti support params")
)

func SetTAG(cli *client.Cli, tag ...string) error {
	if isSupportWTI.Load() == false {
		return ERRNotSupportWTI
	}
	factory.SetTAG(cli, tag...)
	return nil
}

func Update(token ...string) error {
	if isSupportWTI.Load() == false {
		return ERRNotSupportWTI
	}
	factory.Update(token...)
	return nil
}

func BroadCast(content []byte, tag ...string) error {
	if isSupportWTI.Load() == false {
		return ERRNotSupportWTI
	}
	factory.BroadCast(content, tag...)
	return nil
}

func BroadCastByTarget(targetAndContent map[string][]byte) error {
	if isSupportWTI.Load() == false {
		return ERRNotSupportWTI
	}
	factory.BroadCastByTarget(targetAndContent)
	return nil
}

func GetClienterTAGs(token string) ([]string, error) {
	if isSupportWTI.Load() == false {
		return nil, ERRNotSupportWTI
	}
	res := factory.GetClienterTAGs(token)
	return res, nil
}

func GetTAGCreateTime(tag string) (int64, error) {
	if isSupportWTI.Load() == false {
		return 0, ERRNotSupportWTI
	}
	res := factory.GetTAGCreateTime(tag)
	return res, nil
}

func GetTAGClients(tag string) (int64, error) {
	if isSupportWTI.Load() == false {
		return 0, ERRNotSupportWTI
	}
	res := factory.GetTAGClients(tag)
	return res, nil
}

func FlushWTI() error {
	if isSupportWTI.Load() == false {
		return ERRNotSupportWTI
	}
	factory.FlushWTI()
	return nil
}
