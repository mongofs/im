package wti

import (
	"fmt"
	"github.com/mongofs/im"
	"github.com/mongofs/im/client"
	"github.com/mongofs/im/plugins/wti"
	"testing"
)

// 创建IM测试服务器
func Test_IMServer(t *testing.T) {

}

// 这里就是一个Im聊天服务器
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
	tags := []string{"v1","v2"}
	// 调用wti
	wti.Factory.SetTAG(cli.(*client.Cli),tags...)
}

func NewFakeImServer() *Fake {
	res := &Fake{}
	// 设置选项
	options := []im.OptionFunc{
		im.WithBroadCastBuffer(10),
		im.WithServerValidate(res),
	}
	// 设置Option
	option := im.NewOption(options...)
	res.im = im.New(option)
	return res
}


