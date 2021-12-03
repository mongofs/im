package recieve

import (
	"github.com/mongofs/im/client"
)


type Example struct {}


func ( e *Example ) Handle (cli client.Clienter,data []byte){
	cli.Send(data)
}
