#handleReceive

这里received 方法可以在创建客户端 ：basic/router.go ,connection 这个函数中创建客户端
withreceivefunc 这个option中，替代成您自己的接收函数 ，只需要实现统一函数结构receive(cli clienter,data []byte)，
即可注册到客户端中，从客户端socket发送过来的信息您就可以自己处理了