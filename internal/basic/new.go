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
)


func New(opts...  Option)*BasicServer{
	b:= &BasicServer{
		ps:    atomic.Int64{},
		bsIdx: DefaultBucketNumber,
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
		b.rpcPort ="127.0.0.1:8081"
	}

	b.rpc = grpc.NewServer()
	grpc2.RegisterBasicServer(b.rpc,b)
}


func (b * BasicServer)prepareHttpServer(){
	b.http = gin.Default()
	b.initRouter()
}