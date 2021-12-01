package basic

import (
	"golang.org/x/sync/errgroup"
	"log"
	"net"
)



func (s * BasicServer)Run ()error {
	wg := errgroup.Group{}
	wg.Go(s.runhttpServer)
	wg.Go(s.runGrpcServer)
	return wg.Wait()
}


func (s * BasicServer)runGrpcServer ()error{
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


func (s * BasicServer)runhttpServer ()error{
	if s.httpPort ==""{
		s.httpPort ="127.0.0.1:8080"
	}

	log.Println("start http server at %s", s.httpPort)
	if err := s.http.Run(s.httpPort);err !=nil {
		return err
	}
	return nil
}


func (s *BasicServer)Close()error{
	s.rpc.GracefulStop()
	s.cancel()
	return nil
}