package basic

import (
	"context"
	"github.com/gin-gonic/gin"
	grpc2 "github.com/mongofs/api/im/v1"
	"go.uber.org/atomic"
	"google.golang.org/grpc"
	"websocket/internal/basic/bucket"
)

const (
	DefaultBucketSize 	= 1<< 10  //1024
	DefaultBucketNumber = 1<< 7 //128

	DefaultGrpcPort =":8081"
	DefaultHttpPort =":8080"
)


func New(opts...  Option)*BasicServer{
	b:= &BasicServer{
		ps:    atomic.Int64{},
		bsIdx: DefaultBucketNumber,
		rpcPort: DefaultGrpcPort,
		httpPort: DefaultHttpPort,
	}
	b.ps.Store(0)
	for _,o := range opts{
		o(b)
	}
	b.prepareBucketer()
	b.prepareGrpcServer()
	b.prepareHttpServer()

	return b
}



func (b *BasicServer)prepareBucketer(){
	b.bs = make([]bucket.Bucketer,b.bsIdx)
	ctx ,cancel:= context.WithCancel(context.Background())
	for i:= uint32(0);i< b.bsIdx ; i++ {
		b.bs[i] = bucket.New(
			bucket.WithContext(ctx),
			bucket.WithSize(DefaultBucketSize))
	}
	b.cancel=cancel
}



func (b * BasicServer)prepareGrpcServer (){
	if b.rpcPort == ""{
		b.rpcPort =DefaultGrpcPort
	}

	b.rpc = grpc.NewServer()
	grpc2.RegisterBasicServer(b.rpc,b)
}


func (b * BasicServer)prepareHttpServer(){
	b.http = gin.Default()
	b.initRouter()
}