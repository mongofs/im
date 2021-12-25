package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	im "github.com/mongofs/api/im/v1"
)


const (
	waitTime         = 1 << 7
	ProtocolJson     = 1
	ProtocolProtobuf = 2

	MessageTypeText    = 1
	MessageTypeBinary  = 2

)

type client struct {
	lastHeartBeatT int64
	conn           *websocket.Conn
	token          string
	closeFunc      sync.Once
	done           chan struct{}
	ctx            context.Context
	buf            chan []byte
	closeSig       chan<- string
	handleReceive  Receiver

	protocol    int // json /protobuf
	messageType int // text /binary
}

func (c * client)Token()string{
	return c.token
}



func CreateConn(w http.ResponseWriter, r *http.Request,closeSig chan <- string, buffer, messageType, protocol,
						readBuffSize, writeBuffSize int, token string, ctx context.Context,handler Receiver) (Clienter, error) {
	res := &client{
		lastHeartBeatT: time.Now().Unix(),
		done:        make(chan struct{}),
		closeFunc:   sync.Once{},
		buf:         make(chan []byte, buffer),
		token:       token,
		ctx:         ctx,
		closeSig: closeSig,
		protocol:    protocol,
		messageType: messageType,
		handleReceive: handler,
	}
	if err := res.upgrade(w, r, readBuffSize, writeBuffSize); err != nil {
		return nil, err
	}
	if err := res.start(); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *client) upgrade(w http.ResponseWriter, r *http.Request, readerSize, writeSize int) error {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  readerSize,
		WriteBufferSize: writeSize,
	}).Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *client) Send(data []byte, i ...int64) error {
	var (
		sid int64
		d   []byte
		err error
	)
	if len(i) > 0 {
		sid = i[0]
	}
	basic := &im.PushToClient{
		Sid: sid,
		Msg: data,
	}
	if c.protocol == ProtocolJson {
		d, err = json.Marshal(basic)
	} else {
		d, err = proto.Marshal(basic)
	}
	if err != nil {
		return err
	}
	c.send(d)
	return nil
}

func (c *client) LastHeartBeat() int64 {
	return c.lastHeartBeatT
}

func (c *client) send(data []byte) {
	if len(c.buf) *10 > cap(c.buf) * 8 {
		return
	}
	c.buf <- data
}

// param retry ,if retry is ture , don't delete the token
func (c *client) Offline(forRetry ...bool) {
	c.close(forRetry ...)
}

func (c *client) start() error {
	go c.sendProc()
	go c.recvProc()
	return nil
}

func (c *client) sendProc() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(fmt.Sprintf("Client :	 '%v' current panic :'%v'", c.token, err))
		}
	}()
	for {
		select {
		case data := <-c.buf:
			err := c.conn.WriteMessage(c.messageType, data)
			if err != nil {
				// log.Error(err.Error())
				goto loop
			}
		case <-c.done:
			goto loop
		}
	}
loop:
	c.close()
}

// 如果close 是为了重连，就没有
func (c *client) close(forRetry ...bool) {
	flag := false
	if len(forRetry)> 0 {
		flag =forRetry[0]
	}

	c.closeFunc.Do(func() {
		close(c.done)
		c.conn.Close()
		if ! flag {
			c.closeSig <- c.token
		}

		//log.Info(fmt.Sprintf("client : %s is offline",c.token))
	})
}

func (c *client) recvProc() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(fmt.Sprintf("Client :	'%v' current panic :'%v'", c.token, err))
		}
	}()
	for {
		select {
		case <-c.done:
			goto loop
		default:
			_, data, err := c.conn.ReadMessage()
			if err != nil {
				// log.Error(err.Error())
				goto loop
			}
			c.handleReceive.Handle(c,data)
		}
	}
loop:
	c.close()
}


func (c *client) ResetHeartBeatTime(){
	c.lastHeartBeatT =time.Now().Unix()
}


