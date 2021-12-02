package main

import (
	"context"
	im "github.com/mongofs/api/im/v1"
	"testing"
)


// BenchmarkSendMessage-6   	   11966	     93402 ns/op
// 当用户buffer 为1 的情况
func BenchmarkSendMessage(t *testing.B) {
	cli := Client()
	ctx := context.Background()
	data := &im.SendMessageReq{
		Token: "12345",
		Data:  []byte("testdata 1"),
	}
	for i:= 0;i<t.N ;i++ {
		cli.SendMessage(ctx,data)
	}
}



// BenchmarkSendMessageWithBufferSize8-6   	   12408	     97096 ns/op
// 当用户buffer 为8 的情况
func BenchmarkSendMessageWithBufferSize8(t *testing.B) {
	cli := Client()
	ctx := context.Background()
	data := &im.SendMessageReq{
		Token: "12345",
		Data:  []byte("testdata 1"),
	}
	for i:= 0;i<t.N ;i++ {
		cli.SendMessage(ctx,data)
	}
}



// BenchmarkSendMessage-6   	   11626	     94524 ns/op
func BenchmarkBroadCast(t *testing.B) {
	cli := Client()
	ctx := context.Background()
	data := &im.BroadcastReq{
		Data:  []byte("testdata 1"),
	}
	for i:= 0;i<t.N ;i++ {
		cli.Broadcast(ctx,data)
	}
}
