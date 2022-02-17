package im

var VERSION = "master"


func (s *ImSrever)Run ()error {
	var prepareParallelFunc = [] func()error {
		// 启用单独goroutine 进行监控
		s.monitorOnline,
		s.monitorWTI,
		// 启用单独goroutine 进行运行
		s.runGrpcServer,
		s.runhttpServer,
		s.PushBroadCast,
	}
	return ParallelRun(prepareParallelFunc... )
}


// 服务关闭
func (s *ImSrever)Close()error{
	s.rpc.GracefulStop()
	s.cancel()
	return nil
}


