package ack

import (
	"errors"
	"sync"
)


var (
	ErrSidExisted = errors.New("ack : message sid existed ")
)

type ack struct {

	rw sync.RWMutex

	data map[int64] *unit

	retry int8

	retryTime int64

}

const (
	DefaultCapacity = 100
	DefaultRetry =3
	DefaultRetryTime =5
)

func New()Acker {
	return &ack{
		rw:       sync.RWMutex{},
		data:     make(map[int64]*unit,DefaultCapacity),
		retry:   DefaultRetry ,
	}
}


func (a *ack) AddMessage(token string, sid int64, content []byte) error {
	a.rw.Lock()
	defer a.rw.Unlock()

	if _,ok:= a.data[sid];ok {
		return ErrSidExisted
	}
	a.data[sid]= &unit{
		data:  content,
		count: 1,
		token: token,
	}
	return nil
}

func (a *ack) DelMessage(sid int64) {
	a.rw.Lock()
	defer a.rw.Unlock()
	delete(a.data,sid)
}




func (a *ack) Handle(f func(token string, content []byte, sid int64) error) {
	a.rw.Lock()
	defer a.rw.Unlock()
	for sid,v := range a.data {
		if counter := v.addCount();counter > a.retry{
			delete(a.data,sid)
			continue
		}
		if escape := v.escape() ;escape >5 {
			continue
		}
		f(v.token,v.data,sid)
	}
}


