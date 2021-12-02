package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"time"
)
const  (
	Address = "ws://10.0.4.109:8080/conn"
	DefaultRpcAddress = "10.0.4.109:8081"
)


func CreateClient (token string){
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(fmt.Sprintf(Address+"?token=%s",token), nil)
	if nil != err {
		log.Println(err)
		return
	}
	defer conn.Close()
	for {
		messageType, messageData, err := conn.ReadMessage()
		if nil != err {
			log.Println(err)
			break
		}
		switch messageType {
		case websocket.TextMessage://文本数据
			fmt.Println(string(messageData))
		case websocket.BinaryMessage://二进制数据
			fmt.Println(messageData)
		case websocket.CloseMessage://关闭
		case websocket.PingMessage://Ping
		case websocket.PongMessage://Pong
		default:

		}
	}
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}



func GetSliceOfStrings(len int)[]string{
	strs := make([]string,len)
	for i:=0;i<len ; i++{
		strs[i]=RandString(20)
	}
	return strs
}



func main(){
	tokens := GetSliceOfStrings(10000)
	for i:=0;i<10000;i++{
		if i%100 ==0 {
			time.Sleep(1*time.Second)
		}
		go CreateClient(tokens[i])
	}
	select {}
}


