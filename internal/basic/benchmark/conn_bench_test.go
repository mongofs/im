package main

import (
	"math/rand"
	"testing"
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



func GetSliceOfStrings(len int)[]string{
	strs := make([]string,len)
	for i:=0;i<len ; i++{
		strs[i]=RandString(20)
	}
	return strs
}



func TestCreat100Conn(t *testing.T){
	tokens := GetSliceOfStrings(100)
	for _,v :=range tokens{
		go CreateClient(v)
	}
	select {}
}


func TestCreat1000Conn(t *testing.T){
	tokens := GetSliceOfStrings(1000)
	for _,v :=range tokens{
		go CreateClient(v)
	}
	select {}
}


func TestCreat10000Conn(t *testing.T){
	tokens := GetSliceOfStrings(10000)
	for i:=0;i<10000;i++{
		if i%100 ==0 {
			time.Sleep(1*time.Second)
		}
		go CreateClient(tokens[i])
	}
	select {}
}

