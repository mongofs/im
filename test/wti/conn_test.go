package wti

import (
	"fmt"
	"github.com/gorilla/websocket"
	"testing"
	"math/rand"
	"time"
)


var r = rand.New(rand.NewSource(time.Now().Unix()))

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}




func Test_Conn(t *testing.T) {
	tests := []struct{
		tag string
		number int
	}{
		{
			tag: "v1",
			number: 50,
		},
		{
			tag: "v2",
			number: 60,
		},
	}

	for _,v := range tests{
		for i :=0 ;i< v.number;i++ {
			go CreateClient(v.tag)
		}
	}
	time.Sleep(1000 * time.Second)
}


// http://www.baidu.com/conn?token=1080&version=v.10
func CreateClient(version string) {
	token := RandString(20)
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(fmt.Sprintf(Address+"?token=%s&version=%s", token,version), nil)
	if nil != err {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	counter := 0
	/*go func() {
		time.Sleep(50*time.Second)
		conn.WriteMessage(websocket.TextMessage,[]byte(fmt.Sprintf(" heartbeat %s",token)))
	}()*/
	for {
		messageType, messageData, err := conn.ReadMessage()
		if nil != err {
			fmt.Println(err)
			break
		}
		switch messageType {
		case websocket.TextMessage: //文本数据
			counter++
			fmt.Println(token, string(messageData), counter)
		case websocket.BinaryMessage: //二进制数据
		case websocket.CloseMessage: //关闭
		case websocket.PingMessage: //Ping
		case websocket.PongMessage: //Pong
		default:
		}
	}
}
