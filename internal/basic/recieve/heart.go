package recieve

import (
	"websocket/internal/basic/client"
)


type Basic struct {
	// Sid is the unique identifier of the message
	Sid int64
	// MSG is the main body of information transmission
	Msg interface{}
}


func Handle (cli client.Clienter,data []byte){
	/*b := &Basic{}
	json.Unmarshal(data,b)*/
	cli.Send(data)
}
