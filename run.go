package im

import (
	"fmt"
	"net"
	"net/http"
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var VERSION = "master"


func (s *ImSrever)Run ()error {
	wg := errgroup.Group{}
	wg.Go(s.runhttpServer)
	wg.Go(s.runGrpcServer)
	wg.Go(s.monitor)
	return wg.Wait()
}


func (s *ImSrever)runGrpcServer ()error{
	listen, err := net.Listen("tcp", s.rpcPort)
	if err !=nil { log.Fatal(err) }
	log.Info("start GRPC server at ", s.rpcPort)
	if err := s.rpc.Serve(listen);err !=nil {
		log.Fatal(err)
	}

	return nil
}


func (s *ImSrever)runhttpServer ()error{
	listen, err := net.Listen("tcp", s.httpPort)
	if err !=nil { log.Fatal(err) }
	log.Info("start HTTP server at ", s.httpPort)
	if err := http.Serve(listen,s.http);err !=nil {
		log.Fatal(err)
	}
	return nil
}


func (s *ImSrever)Close()error{
	s.rpc.GracefulStop()
	s.cancel()
	return nil
}


func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf(" %s:%d", filename, f.Line)
		},
	})
}