package main

import (
	"testing"
	"time"
)



// 测试100 个用户同时在线
func TestCreat100Conn(t *testing.T){
	tokens := GetSliceOfStrings(100)
	for _,v :=range tokens{
		go CreateClient(v)
	}
	select {}
}

// 测试3000 个用户同时在线
func TestCreat3000Conn(t *testing.T){
	tokens := GetSliceOfStrings(3000)
	for _,v :=range tokens{
		time.Sleep(20*time.Millisecond)
		go CreateClient(v)
	}

	select {}
}

// 测试10000 个用户同时在线
func TestCreat10000Conn(t *testing.T){
	tokens := GetSliceOfStrings(10000)
	for i:=0;i<100000;i++{
		if i%100 ==0 {
			time.Sleep(1*time.Second)
		}
		go CreateClient(tokens[i])
	}
	select {}
}

