package wti

import (
	"context"
	"fmt"
	im "github.com/mongofs/api/im/v1"
	"google.golang.org/grpc"
	"testing"
	"time"
)


const (
	Address           = "ws://127.0.0.1:8080/conn"
	DefaultRpcAddress = "127.0.0.1:8081"
)

var (
	cli = Client()
	ctx = context.Background()
)

func Client ()im.BasicClient{
	return im.NewBasicClient(conn)
}

var conn,_ = grpc.Dial(DefaultRpcAddress,grpc.WithInsecure())

func TestSendMessage(t *testing.T) {
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
		fmt.Println(cli.WTIBroadcast(ctx,v,))
	}
	// output
	// nil
	// nil
	// nil
	// nil
}

func TestDistribute(t *testing.T) {
	for {
		distribute,err :=cli.WTIDistribute(ctx,&im.Empty{})
		if err !=nil {
			t.Fatal(err)
		}
		fmt.Printf("当前用户分布： %+v\n\r ",distribute)
		time.Sleep(5*time.Second)
	}
}

