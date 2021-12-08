package im

import (
	"github.com/mongofs/im/bucket"
	"go.uber.org/atomic"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type ImSrever struct {
	http     *http.ServeMux
	rpc      *grpc.Server
	bs       []bucket.Bucketer
	ps       atomic.Int64
	cancel   func()

	opt *Option
}




func (s *ImSrever)monitor ()error{
	for{
		n := int64(0)
		for _,bck := range  s.bs{
			bck.Flush()
			n += bck.Onlines()
		}
		s.ps.Store(n)
		time.Sleep(10 *time.Second)
	}
	return nil
}


func (s *ImSrever) bucket(token string) bucket.Bucketer {
	idx := Index(token,uint32(s.opt.ServerBucketNumber))
	return s.bs[idx]
}





