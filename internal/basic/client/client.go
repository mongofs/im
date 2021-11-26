package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var (
	ErrUpgradeBadConn =errors.New("client: conn reader/writer is bad ")
	ErrConnCreatedError =errors.New("client: create connection is failed ")
	ErrNoSupplyUserToken =errors.New("client: create connection is failed ")
	ErrContextNotSupply =errors.New("client: context is not supply ")
)

type Sender  interface {
	Send([]byte)
	SendWithAck([]byte)
}

// 	client 结构体是为了保存单个用户连接信息内容，包含
//	conn 用户连接地址
//	usertoken
//	ctx 控制goroutine 及时退出的上下文
//  dataQueue  用户信息缓冲区
//  用户信息缓存区
type client struct {
	// writer
	writer http.ResponseWriter
	// reader
	reader *http.Request
	// 用户连接地址
	conn *websocket.Conn
	// user token should be validate
	token string
	// close flag
	closeFunc  sync.Once
	// 通知用户的goroutine 退出
	done       chan struct{}
	// 控制读写goroutine 退出的ctx
	ctx context.Context
	// 用户信息缓冲区
	buffer chan []byte // length should be 1
	// close sig
	closeSig 	chan<- string
}


const  (
	DefaultACKLength =8
)

func New(opt...optionFunc)(*client,error) {
	res := &client{}
	for _,o := range opt {
		o(res)
	}

	if err := res.validate();err!=nil{
		return nil, err
	}
	if err := res.Upgrade() ;err !=nil {return nil, err}
	return res,nil
}

func (c *client)Upgrade()error{
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(c.writer, c.reader, nil)
	if err !=nil { return err}
	c.conn =conn
	return nil
}


func (c *client) Start(ctx context.Context) error{
	go c.sendProc()
	go c.recvProc()
	return nil
}


func (c *client) sendProc() {
	defer func() {
		if err :=recover();err !=nil {
			log.Error(fmt.Sprintf("Client :	 '%v' current panic :'%v'",c.token,err))
		}
	}()
	for {
		select {
		case data := <-c.buffer:
			err := c.conn.WriteMessage(websocket.TextMessage,data)
			if err != nil {
				if err==websocket.ErrCloseSent{
					log.Info(fmt.Sprintf("Cliet : '%v' soket conn is break , reason : %v " , c.token, err ) )
					c.close()
				}else {
					log.Info(fmt.Sprintf("Cliet : '%v' soket conn is break , reason : %v " , c.token, err ) )
					continue
				}
			}
		case <-c.done:
			log.Info(fmt.Sprintf("Cliet : '%v' sender goroutine is close " , c.token) )
			break
		}
	}
}


func (c *client) close() {
	c.closeFunc.Do(func() {
		close(c.done)
		c.closeSig<-c.token
		c.conn.Close()
	})
}


// 接收到的消息进行
func (c *client) recvProc() {

	defer func() {
		if err :=recover();err !=nil {
			log.Error(fmt.Sprintf("Client :	'%v' current panic :'%v'",c.token,err))
		}
	}()

	for {
		select {
		case <-c.done:
			log.Info(fmt.Sprintf("Client : '%v' reciver quite safely" , c.token) )
			break
		default:
			_, data, err := c.conn.ReadMessage()
			if err != nil {
				log.Error(fmt.Sprintf("Cliet : '%v' read soketconn current error , reason : %v " , c.token, err ) )
				break
			}
			// 将dis
			c.handleRecv(data)
		}
	}
	c.close()
}






func (c *client) validate ()error {
	if c.token =="" {
		return ErrNoSupplyUserToken
	}
	if c.ctx == nil {
		return ErrContextNotSupply
	}
	if c.writer == nil {
		return ErrUpgradeBadConn
	}
	if c.reader == nil {
		return ErrUpgradeBadConn
	}
	return nil
}




