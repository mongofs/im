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

	push:= `{"id":1041584,"user_id":1041584,"cmd":1003,"message_content":"{\"id\":2574952,
			\"match_id\":1041584,\"type\":0,\"team_main\":\"巴拉卡斯中央\",\"team_cust\":\"基尔梅斯\",
			\"main_scale\":\"0\",\"guest_scale\":\"0\",\"team_event\":1,
			\"user_event\":\"费德里科安塞尔莫\"}"}`


	data := &im.SendMessageReq{
		Token: "IYURAPAURILTIUZBJBDT",
		Data:  []byte(push),
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



//BenchmarkBroadCast-6   	    1932	    622000 ns/op 串行
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


// BenchmarkBroadCast_paraml-6   	    2751	    391783 ns/op
// BenchmarkBroadCast_paraml-6   	    3105	    494931 ns/op
func BenchmarkBroadCast_paraml(t *testing.B) {
	cli := Client()
	ctx := context.Background()
	data := &im.BroadcastReq{
		Data:  []byte("testdata 1"),
	}
	for i:= 0;i<t.N ;i++ {
		cli.Broadcast(ctx,data)
	}
}


// 压测分类广播，大致情况如下
// BenchmarkBroadCastByWTI-6   	    2905	    411497 ns/op
// BenchmarkBroadCastByWTI-6   	    3099	    421722 ns/op
// BenchmarkBroadCastByWTI-6   	    2709	    385287 ns/op
func BenchmarkBroadCastByWTI(t *testing.B){
	cli := Client()
	ctx := context.Background()
	tests:= &im.BroadcastByWTIReq{Data: map[string][]byte{
		"v1":[]byte("fucker you v2"),
		"v2":[]byte("lover you v1"),
		"v3":[]byte("what are you v1"),
	}}

	for i:= 0;i<t.N ;i++ {
		cli.BroadcastByWTI(ctx,tests,)
	}
}