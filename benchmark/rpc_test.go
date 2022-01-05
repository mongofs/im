//To test im, it is recommended to use the following website: http://coolaf.com/tool/chattest ,
//the connection method is directly used: WS: // 127.0.0.1:8080/conn?Token = 12345. Establish a
//connection with the IM server, use the function, and run the test method


// It is recommended to test im using the following URL: http://coolaf.com/tool/chattest

// First run the startup method under CMD to start the logic server of IM
// Step 2: open the im test website to test the connection: ws://127.0.0.1:8080/conn?token=12345
// Step 3 run the testsendmessage method

package main

import (
	"context"
	"fmt"
	im "github.com/mongofs/api/im/v1"
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	cli := Client()
	ctx := context.Background()

	tests:= []*im.SendMessageReq{
		&im.SendMessageReq{
			Token: "12345",
			Data:  []byte("testdata 1"),
		},
		&im.SendMessageReq{
			Token: "12345",
			Data:  []byte("testdata 2"),
		},
		&im.SendMessageReq{
			Token: "12345",
			Data:  []byte("testdata 3"),
		},
		&im.SendMessageReq{
			Token: "12345",
			Data:  []byte("testdata 4"),
		},
	}

	for _,v := range tests {
		fmt.Println(cli.SendMessage(ctx,v))
	}

	// output
	// nil
	// nil
	// nil
	// nil
}


func TestBroadCast(t *testing.T){
	cli := Client()
	ctx := context.Background()

	tests:= []*im.BroadcastReq{
		&im.BroadcastReq{
			Data:  []byte("broadcast 1"),
		},
		&im.BroadcastReq{
			Data:  []byte("broadcast 2"),
		},
		&im.BroadcastReq{
			Data:  []byte("broadcast 3"),
		},
		&im.BroadcastReq{
			Data:  []byte("broadcast 4"),
		},
	}

	for _,v := range tests {
		fmt.Println(cli.Broadcast(ctx,v))
	}
	// output
	// nil
	// nil
	// nil
	// nil
}


func TestOnlines(t *testing.T){
	cli := Client()
	ctx := context.Background()
	var n =0
	tk :=time.NewTicker(1*time.Second)
	for {
		<- tk.C
		fmt.Println(cli.Onliens(ctx,&im.Empty{}))
		n += 1
		if n ==10 {
			break
		}
	}
	// output
	// nil
	// nil
	// nil
	// nil
}



func TestTickerBroadCast(t *testing.T){
	cli := Client()
	ctx := context.Background()
	var push = `{"id":1041584,"user_id":1041584,"cmd":1003,"message_content":"{\"id\":2574952,
			\"match_id\":1041584,\"type\":0,\"team_main\":\"巴拉卡斯中央\",\"team_cust\":\"基尔梅斯\",
			\"main_scale\":\"0\",\"guest_scale\":\"0\",\"team_event\":1,
			\"user_event\":\"费德里科安塞尔莫\"}"}`

	tests:= 	&im.BroadcastReq{
		Data:  []byte(push),
	}


	counter := 0
	// 50 下推
	for {
		time.Sleep(100*time.Millisecond)
		res,err := cli.Broadcast(ctx,tests)
		counter++
		if err !=nil {
			fmt.Println("error current :",err,counter)
			continue
		}
		fmt.Println("当前堆积：",res.Size,",当前消息序号",counter)
	}
}


