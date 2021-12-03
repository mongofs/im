package im

import (
	"golang.org/x/sync/errgroup"
	"log"
	"net"
)



func (s *ImSrever)Run ()error {
	wg := errgroup.Group{}
	wg.Go(s.runhttpServer)
	wg.Go(s.runGrpcServer)
	wg.Go(s.monitor)
	return wg.Wait()
}


func (s *ImSrever)runGrpcServer ()error{
	listen, err := net.Listen("tcp", s.rpcPort)
	if err !=nil {
		log.Fatal(err)
	}
	log.Println("start grpc server at %s", s.rpcPort)
	if err := s.rpc.Serve(listen);err !=nil {
		return err
	}
	return nil
}


func (s *ImSrever)runhttpServer ()error{
	if s.httpPort ==""{
		s.httpPort ="127.0.0.1:8080"
	}

	log.Println("start http server at %s", s.httpPort)
	if err := s.http.Run(s.httpPort);err !=nil {
		return err
	}
	return nil
}


func (s *ImSrever)Close()error{
	s.rpc.GracefulStop()
	s.cancel()
	return nil
}