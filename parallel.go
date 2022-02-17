package im

import (
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"time"
	"github.com/mongofs/im/plugins/wti"
)


func ParallelRun (parallels ... func()error)error{
	wg := errgroup.Group{}
	for _,v:= range parallels {
		wg.Go(v)
	}
	return wg.Wait()
}


// 统计用户在线人数
// 监控buffer 长度 并进行报警
func (s *ImSrever) monitorOnline() error {
	for {
		n := int64(0)
		for _, bck := range s.bs {
			bck.Flush()
			n += bck.Onlines()
		}
		s.ps.Store(n)
		time.Sleep(10 * time.Second)
	}
	return nil
}

// 统计用户在线人数
// 监控buffer 长度 并进行报警
func (s *ImSrever) monitorWTI() error {
	if s.opt.SupportPluginWTI {
		for {
			wti.FlushWTI()
			time.Sleep(20*time.Second)
		}
	}
	return nil

}

// 监控rpc 服务
func (s *ImSrever)runGrpcServer ()error{
	listen, err := net.Listen("tcp", s.opt.ServerRpcPort)
	if err !=nil { s.opt.ServerLogger.Fatal(err) }
	s.opt.ServerLogger.Infof("im/run : start GRPC server at %s ", s.opt.ServerRpcPort)
	if err := s.rpc.Serve(listen);err !=nil {
		s.opt.ServerLogger.Fatal(err)
	}
	return nil
}

// 监控http服务
func (s *ImSrever)runhttpServer ()error{
	listen, err := net.Listen("tcp", s.opt.ServerHttpPort)
	if err !=nil { s.opt.ServerLogger.Fatal(err) }
	s.opt.ServerLogger.Infof("im/run : start HTTP server at %s ", s.opt.ServerHttpPort)
	if err := http.Serve(listen,s.http);err !=nil {
		s.opt.ServerLogger.Fatal(err)
	}
	return nil
}



// 单独处理广播业务
func (s *ImSrever) PushBroadCast() error {
	wg := errgroup.Group{}
	for i := 0; i < s.opt.BroadCastHandler; i++ {
		wg.Go(func() error {
			for {
				req := <-s.buffer
				for _, v := range s.bs {
					err := v.BroadCast(req.Data, false)
					if err != nil {
						s.opt.ServerLogger.Error(err)
					}
				}
			}
			return nil
		})
	}
	return wg.Wait()
}