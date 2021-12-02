package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"

)

const  (
	Address = "ws://127.0.0.1:8080/conn"
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


