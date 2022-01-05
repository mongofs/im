package im

import (
	"net"
	"net/http"
	"golang.org/x/sync/errgroup"
)

var VERSION = "master"


func (s *ImSrever)Run ()error {
	wg := errgroup.Group{}
	wg.Go(s.runhttpServer)
	wg.Go(s.runGrpcServer)
	wg.Go(s.monitor)
	wg.Go(s.PushBroadCast)
	return wg.Wait()
}


func (s *ImSrever)runGrpcServer ()error{
	listen, err := net.Listen("tcp", s.opt.ServerRpcPort)
	if err !=nil { s.opt.ServerLogger.Fatal(err) }
	s.opt.ServerLogger.Infof("im/run : start GRPC server at %s ", s.opt.ServerRpcPort)
	if err := s.rpc.Serve(listen);err !=nil {
		s.opt.ServerLogger.Fatal(err)
	}

	return nil
}


func (s *ImSrever)runhttpServer ()error{
	listen, err := net.Listen("tcp", s.opt.ServerHttpPort)
	if err !=nil { s.opt.ServerLogger.Fatal(err) }
	s.opt.ServerLogger.Infof("im/run : start HTTP server at %s ", s.opt.ServerHttpPort)
	if err := http.Serve(listen,s.http);err !=nil {
		s.opt.ServerLogger.Fatal(err)
	}
	return nil
}


func (s *ImSrever)Close()error{
	s.rpc.GracefulStop()
	s.cancel()
	return nil
}


