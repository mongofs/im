package recieve

import "websocket/client"

type Receiver interface {

	Handle (cli client.Clienter,data []byte)
}