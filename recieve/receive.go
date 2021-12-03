package recieve

import "github.com/mongofs/im/client"

type Receiver interface {

	Handle (cli client.Clienter,data []byte)
}