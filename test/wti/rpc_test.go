package wti

import (
	"context"
	"fmt"
	im "github.com/mongofs/api/im/v1"
	"google.golang.org/grpc"
	"testing"
)


const (
	Address           = "ws://127.0.0.1:8080/conn"
	DefaultRpcAddress = "127.0.0.1:8081"
)


func TestSendMessage(t *testing.T) {
	cli := Client()
	ctx := context.Background()
	tests:= []*im.BroadcastByWTIReq{
		{
			Data: map[string][]byte{
				"v1":[]byte("fucker you v2"),
				"v2":[]byte("lover you v1"),
				"v3":[]byte("what are you v1"),
			},
		},
	}
	for _,v := range tests {
		fmt.Println(cli.BroadcastByWTI(ctx,v,))
	}
	// output
	// nil
	// nil
	// nil
	// nil
}

var conn,_ = grpc.Dial(DefaultRpcAddress,grpc.WithInsecure())


func Client ()im.BasicClient{
	return im.NewBasicClient(conn)
}
