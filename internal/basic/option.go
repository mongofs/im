package basic

type Option func (b *BasicServer)


func WithHttpPort (httpPort string)Option{
	return func(b *BasicServer) {
		b.httpPort=httpPort
	}
}



func WithRpcPort (Port string)Option{
	return func(b *BasicServer) {
		b.rpcPort=Port
	}
}