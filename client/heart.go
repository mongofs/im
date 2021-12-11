package client

type Example struct {}


func ( e *Example) Handle (cli Clienter,data []byte){
	cli.Send(data)
}
