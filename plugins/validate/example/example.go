package example

import (
	"errors"
	"fmt"
	"github.com/mongofs/im/client"
)

type  DefaultValidate struct {
}


func (d *DefaultValidate) Validate(token string)error{
	if token == "" {
		return errors.New("token is not good ")
	}
	return nil
}





func (d *DefaultValidate)ValidateFailed(err error,cli client.Clienter){

	fmt.Println(err.Error())
	// 当用户登录验证失败，逻辑应该在这里来处理
	cli.Send([]byte("user validate is bad"))
	cli.Offline()
}


func (d *DefaultValidate)ValidateSuccess(cli client.Clienter){
	// 当用户登录验证失败，逻辑应该在这里来处理
	cli.Send([]byte("user is online "))
}
