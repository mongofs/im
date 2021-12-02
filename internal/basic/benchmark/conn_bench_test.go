package main

import (
	"testing"
	"time"
)




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


func TestCreat100000Conn(t *testing.T){
	tokens := GetSliceOfStrings(100000)
	for i:=0;i<100000;i++{
		if i%100 ==0 {
			time.Sleep(1*time.Second)
		}
		go CreateClient(tokens[i])
	}
	select {}
}

