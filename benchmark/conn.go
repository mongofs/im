package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"time"
)



func CreateClient (token string){
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(fmt.Sprintf(Address+"?token=%s",token), nil)
	if nil != err {
		log.Println(err)
		return
	}
	defer conn.Close()

	counter :=0

	go func() {
		time.Sleep(10*time.Second)
		conn.WriteMessage(websocket.TextMessage,[]byte(fmt.Sprintf(" heartbeat %s",token)))
	}()

	for {
		messageType, messageData, err := conn.ReadMessage()
		if nil != err {
			log.Println(err)
			break
		}
		switch messageType {
		case websocket.TextMessage://文本数据
			counter++
			if token == "1234"{
				fmt.Println(string(messageData),"这是消息：",counter)
			}
		case websocket.BinaryMessage://二进制数据

		case websocket.CloseMessage://关闭
		case websocket.PingMessage://Ping
		case websocket.PongMessage://Pong
		default:

		}
	}
}


// 用户断线重连
func CreateClientAndTickerReConn (token string){
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(fmt.Sprintf(Address+"?token=%s",token), nil)
	if nil != err {
		log.Println(err)
		return
	}
	defer conn.Close()

	counter :=0

	for {
		messageType, messageData, err := conn.ReadMessage()
		if nil != err {
			log.Println(err)
			break
		}
		switch messageType {
		case websocket.TextMessage://文本数据
			counter++
			if token == "1234"{
				fmt.Println(string(messageData),"这是消息：",counter)
			}
		case websocket.BinaryMessage://二进制数据

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
		strs[i]= RandString(20)
	}
	return strs
}




