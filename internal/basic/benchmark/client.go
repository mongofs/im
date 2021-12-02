//To test websocket, it is recommended to use the following website: http://coolaf.com/tool/chattest ,
//the connection method is directly used: WS: // 127.0.0.1:8080/conn?Token = 12345. Establish a
//connection with the IM server, use the function, and run the test method

package main

import (
	im "github.com/mongofs/api/im/v1"
	"google.golang.org/grpc"
)


var conn,_ = grpc.Dial("127.0.0.1:8081",grpc.WithInsecure())


func Client ()im.BasicClient{
	return im.NewBasicClient(conn)
}
