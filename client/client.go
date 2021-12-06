package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	im "github.com/mongofs/api/im/v1"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

var (
	ErrUpgradeBadConn =errors.New("client: conn reader/writer is bad ")
	ErrhandReceiveIsNil =errors.New("client: handler receive is nil ")
	ErrConnCreatedError =errors.New("client: create connection is failed ")
	ErrNoSupplyUserToken =errors.New("client: create connection is failed ")
	ErrContextNotSupply =errors.New("client: context is not supply ")
)



type client struct {
	writer http.ResponseWriter
	reader *http.Request
	conn *websocket.Conn
	token string
	closeFunc  sync.Once
	done       chan struct{}
	ctx context.Context
	buffer chan []byte
	closeSig 	chan<- string
	handleReceive func(cli Clienter,data []byte)
	agreement int
}

func (c *client) Send(data []byte, i ...int64)error {
	var (
		sid int64
		d []byte
		err error
	)

	if len(i) >0 {
		sid = i[0]
	}

	basic := &im.PushToClient{
		Sid: sid,
		Msg: data,
	}
	if c.agreement == AgreementJson {
		d,err = json.Marshal(basic)
	}else {
		d,err = proto.Marshal(basic)
	}
	if err!=nil{
		return err
	}
	c.send(d)
	return nil
}


func (c *client) send(data []byte) {
	c.buffer<- data
}



func (c *client) Offline() {
	c.close()
}



func New(opt...OptionFunc)(Clienter,error) {
	res := &client{
		buffer: make(chan []byte,10),
		done: make(chan struct{}),
		closeFunc: sync.Once{},
		agreement: AgreementJson,
	}
	for _,o := range opt {
		o(res)
	}
	if err := res.validate();err!=nil{
		return nil, err
	}
	if err := res.upgrade() ;err !=nil {return nil, err}
	if err := res.start();err !=nil {return nil, err}
	return res,nil
}

func (c *client)upgrade()error{
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(c.writer, c.reader, nil)
	if err !=nil { return err}
	c.conn =conn
	return nil
}


func (c *client) start() error{
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
					goto loop
				}else {
					log.Info(fmt.Sprintf("Cliet : '%v' soket conn is break , reason : %v " , c.token, err ) )
					continue
				}
			}
		case <-c.done:
			goto loop
		}
	}
	loop :
		log.Info(fmt.Sprintf("Cliet : '%v' sender goroutine is close " , c.token) )
}


const (
	waitTime = 1 <<7
)

func (c *client) close() {
	c.closeFunc.Do(func() {
		close(c.done)
		time.Sleep(waitTime * time.Millisecond)
		c.conn.Close()
		c.closeSig<-c.token
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
			goto loop
		default:
			_, data, err := c.conn.ReadMessage()
			if err !=nil {
				log.Error(fmt.Sprintf("Cliet : '%v' read soketconn current error , reason : %v " , c.token, err ) )
				goto loop
			}
			c.handleReceive(c, data)
		}
	}
	loop:
		log.Info(fmt.Sprintf("Client : '%v' reciver quite safely" , c.token) )
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
	if c.handleReceive == nil {
		return ErrhandReceiveIsNil
	}

	return nil
}




