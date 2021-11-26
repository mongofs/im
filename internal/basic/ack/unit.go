package ack

import "time"

type unit struct {
	data []byte
	count int8
	token string
	sendTime time.Duration
}


func (u *unit)addCount()int8{
	u.count++
	return u.count
}


func (u *unit)escape ()int64{
	return 0
}