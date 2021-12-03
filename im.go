package im

import (
	"context"
	"github.com/gin-gonic/gin"
	grpc2 "github.com/mongofs/api/im/v1"
	"go.uber.org/atomic"
	"google.golang.org/grpc"
	"websocket/bucket"
	"websocket/recieve"
	"websocket/validate"
	"websocket/validate/example"
)

const (
	DefaultBucketSize 	= 1<< 10  //1024
	DefaultBucketNumber = 1<< 6 //64

	DefaultGrpcPort =":8081"
	DefaultHttpPort =":8080"
)


var (
	DefaultValidate  validate.Validater = &example.DefaultValidate{}
	DefaultReciever  recieve.Receiver = &recieve.Example{}
 )


func New(opts...Option)*ImSrever {
	b:= &ImSrever{
		ps:       atomic.Int64{},
		bsIdx:    DefaultBucketNumber,
		rpcPort:  DefaultGrpcPort,
		httpPort: DefaultHttpPort,
		recevier: DefaultReciever,
		validate: DefaultValidate,
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



func (b *ImSrever)prepareBucketer(){
	b.bs = make([]bucket.Bucketer,b.bsIdx)
	ctx ,cancel:= context.WithCancel(context.Background())
	for i:= uint32(0);i< b.bsIdx ; i++ {
		b.bs[i] = bucket.New(
			bucket.WithContext(ctx),
			bucket.WithSize(DefaultBucketSize))
	}
	b.cancel=cancel
}



func (b *ImSrever)prepareGrpcServer (){
	b.rpc = grpc.NewServer()
	grpc2.RegisterBasicServer(b.rpc,b)
}


func (b *ImSrever)prepareHttpServer(){
	b.http = gin.Default()
	b.initRouter()
}