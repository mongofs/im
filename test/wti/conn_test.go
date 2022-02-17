package wti

import (
	"fmt"
	"github.com/gorilla/websocket"
	"testing"
	"time"
)

// 模拟创建3个链接： 分别是订阅
// conn1 : v1
// conn2 : v2
// conn3 : v1,v2


func Test_Conn(t *testing.T){
	go CreateClient("v1")
	go CreateClient("v2")
	go CreateClient("v3")
	//CreateClient("v1&v2")
	time.Sleep(1000 *time.Second)
}

func CreateClient (token string){
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(fmt.Sprintf(Address+"?token=%s&ver=v1",token), nil)
	if nil != err {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	counter :=0

	go func() {
		time.Sleep(50*time.Second)
		conn.WriteMessage(websocket.TextMessage,[]byte(fmt.Sprintf(" heartbeat %s",token)))
	}()

	for {
		messageType, messageData, err := conn.ReadMessage()
		if nil != err {
			fmt.Println(err)
			break
		}
		switch messageType {
		case websocket.TextMessage://文本数据
			counter++
			fmt.Println(token,string(messageData),counter)
		case websocket.BinaryMessage://二进制数据
		case websocket.CloseMessage://关闭
		case websocket.PingMessage://Ping
		case websocket.PongMessage://Pong
		default:
		}
	}
}