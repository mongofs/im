package im

import (
	"context"
	grpc2 "github.com/mongofs/api/im/v1"
	"github.com/mongofs/im/bucket"
	"go.uber.org/atomic"
	"google.golang.org/grpc"
	"net/http"
)


func New(opts *Option) *ImSrever {
	b := &ImSrever{
		ps:     atomic.Int64{},
		opt:    opts,
	}
	b.buffer = make(chan *grpc2.BroadcastReq,opts.ServerBuffer)
	b.ps.Store(0)
	b.prepareBucketer()
	b.prepareGrpcServer()
	b.prepareHttpServer()
	return b
}

func (h *ImSrever) prepareBucketer() {
	h.bs = make([]bucket.Bucketer, h.opt.ServerBucketNumber)
	_, cancel := context.WithCancel(context.Background())
	h.cancel = cancel

	BucketOptionSet := &bucket.Option{
		HeartBeatInterval: int64(h.opt.ClientHeartBeatInterval),
		ReaderBufferSize:  h.opt.ClientReaderBufferSize,
		WriteBufferSize:   h.opt.ClientWriteBufferSize,
		ClientBufferSize:  h.opt.ClientBufferSize,
		MessageType:       h.opt.ClientMessageType,
		Protocol:          h.opt.ClientProtocol,
		BucketSize:        h.opt.BucketSize,
	}

	for i:= 0 ;i<h.opt.ServerBucketNumber;i ++ {
		h.bs[i] = bucket.New(BucketOptionSet)
	}
}

func (b *ImSrever) prepareGrpcServer() {
	b.rpc = grpc.NewServer()
	grpc2.RegisterBasicServer(b.rpc, b)
}

func (b *ImSrever) prepareHttpServer() {
	b.http = http.NewServeMux()
	b.initRouter()
}


