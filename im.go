package im

import (
	"context"
	grpc2 "github.com/mongofs/api/im/v1"
	"github.com/mongofs/im/bucket"
	"github.com/mongofs/im/recieve"
	"github.com/mongofs/im/validate"
	"github.com/mongofs/im/validate/example"
	"go.uber.org/atomic"
	"google.golang.org/grpc"
	"net/http"
)

const (
	DefaultBucketSize   = 1 << 10 //1024
	DefaultBucketNumber = 1 << 5  //32

	DefaultGrpcPort          = ":8081"
	DefaultHttpPort          = ":8080"
	DefaultMessageSendMethod = 1 // text / byte
	DefaultAgreement         = 1 //json  /proto
)

var (
	DefaultValidate validate.Validater = &example.DefaultValidate{}
	DefaultReciever recieve.Receiver   = &recieve.Example{}
)

func New(opts ...Option) *ImSrever {
	b := &ImSrever{
		ps:                atomic.Int64{},
		bsIdx:             DefaultBucketNumber,
		rpcPort:           DefaultGrpcPort,
		httpPort:          DefaultHttpPort,
		MessageSendMethod: DefaultMessageSendMethod,
		agreement:         DefaultAgreement,
		recevier:          DefaultReciever,
		validate:          DefaultValidate,
	}
	b.ps.Store(0)
	for _, o := range opts {
		o(b)
	}
	b.prepareBucketer()
	b.prepareGrpcServer()
	b.prepareHttpServer()
	return b
}

func (b *ImSrever) prepareBucketer() {
	b.bs = make([]bucket.Bucketer, b.bsIdx)
	ctx, cancel := context.WithCancel(context.Background())
	for i := uint32(0); i < b.bsIdx; i++ {
		b.bs[i] = bucket.New(
			bucket.WithContext(ctx),
			bucket.WithSize(DefaultBucketSize))
	}
	b.cancel = cancel
}

func (b *ImSrever) prepareGrpcServer() {
	b.rpc = grpc.NewServer()
	grpc2.RegisterBasicServer(b.rpc, b)
}

func (b *ImSrever) prepareHttpServer() {
	b.http = http.NewServeMux()
	b.initRouter()
}
