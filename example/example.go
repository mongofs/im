package main
import (
	"log"

	"github.com/mongofs/im"
)


func main(){
	serv := NewServer()
	if err :=serv.Run();err  !=nil {
		log.Fatal(err)
	}
}

type simpleServer struct {
	im *im.ImSrever
}

func NewServer ()*simpleServer{
	imServ:= im.New(im.DefaultOption())
	return &simpleServer{
		im: imServ,
	}
}

func (s *simpleServer) Run()error{
	return s.im.Run()
}


func (s *simpleServer)Close()error{
	s.im.Close()
	return nil
}

