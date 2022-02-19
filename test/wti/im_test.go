package wti

import (
	"fmt"
	"github.com/mongofs/im"
	"github.com/mongofs/im/client"
	"github.com/mongofs/im/plugins/wti"
	"testing"
)

// http://www.baidu.com/conn?token=1080&version=v.10

// 创建IM测试服务器，模拟测试im服务器的具体内容。
func Test_IMServer(t *testing.T) {
	wti.SetSupport()
	serv := NewFakeImServer()
	serv.Run()
}

// this is a temp imserver
type Fake struct {
	im *im.ImSrever
}
func (f *Fake) Run (){
	err := f.im.Run()
	if err != nil {
		fmt.Println(err)
	}
	f.im.Close()
}

func (f *Fake)Validate(token string)error{
	return nil
}
func (f *Fake)ValidateFailed(err error,cli client.Clienter){
	fmt.Println(err)
}
func (f *Fake)ValidateSuccess(cli client.Clienter){
	// 可以通过header 或者 get query 方式来传参，或者从数据库获取当前用户的tag
	req := cli.Request()
	res := req.Form["version"]
	tags := []string{res[0]}
	// 调用wti
	if err := wti.SetTAG(cli.(*client.Cli),tags...);err != nil {
		fmt.Println(err)
	}
}

func NewFakeImServer() *Fake {
	res := &Fake{}
	// 设置选项
	options := []im.OptionFunc{
		im.WithBroadCastBuffer(10),
		im.WithServerValidate(res),
		im.WithPluginsWTI(true),
	}
	// 设置Option
	option := im.NewOption(options...)
	res.im = im.New(option)
	return res
}


