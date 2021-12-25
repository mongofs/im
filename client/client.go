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

type Cli struct {
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

func (c * Cli)Token()string{
	return c.token
}



func CreateConn(w http.ResponseWriter, r *http.Request,closeSig chan <- string, buffer, messageType, protocol,
						readBuffSize, writeBuffSize int, token string, ctx context.Context,handler Receiver) (Clienter, error) {
	res := &Cli{
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

func (c *Cli) upgrade(w http.ResponseWriter, r *http.Request, readerSize, writeSize int) error {
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

func (c *Cli) Send(data []byte, i ...int64) error {
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

func (c *Cli) LastHeartBeat() int64 {
	return c.lastHeartBeatT
}

func (c *Cli) send(data []byte) {
	if len(c.buf) *10 > cap(c.buf) * 8 {
		return
	}
	c.buf <- data
}

// param retry ,if retry is ture , don't delete the token
func (c *Cli) Offline() {
	c.close(false)
}


func (c *Cli)OfflineForRetry(retry ...bool){
	c.close(retry...)
}


func (c *Cli) start() error {
	go c.sendProc()
	go c.recvProc()
	return nil
}

func (c *Cli) sendProc() {
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
func (c *Cli) close(forRetry ...bool) {
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

func (c *Cli) recvProc() {
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


func (c *Cli) ResetHeartBeatTime(){
	c.lastHeartBeatT =time.Now().Unix()
}


