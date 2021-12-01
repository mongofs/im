package recieve

import (
	"websocket/internal/basic/client"
)

func Handle (cli client.Clienter,data []byte){

	cli.Send(data)
}
