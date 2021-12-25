package main

import "testing"

// 测试重复连接的情况



func TestRepeatConnection(t *testing.T){

	var token = "12345"

	for i:= 0; i<2;i++ {
		go CreateClient(token)
	}
	select {}
}
