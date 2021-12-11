package client

type Receiver interface {

	Handle (cli Clienter,data []byte)
}