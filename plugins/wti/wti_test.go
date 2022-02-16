package wti

import (
	"fmt"
	"testing"
)

// 广播，针对tag进行广播，这也是wti的核心接口，分类广播也是基于这个接口
func TestTg_BroadCast(t *testing.T) {
	tests := []struct {
		target []string
		content []byte
	}{
		{
			target: []string{"v1"} ,
			content: []byte("hello content"),
		},
		{
			target: []string{"v1","v2"} ,
			content: []byte("hello content"),
		},
		{
			target: []string{"v1","v2","v3"} ,
			content: []byte("hello content"),
		},
	}
	// 测试 v1、v2 、v1 三个版本的不同，如果要模式真实连接，则需要执行
	// 1. 创建im 连接服务器，并开启wti 配置参数
	// 2. 要在handler方法中调用factory 进行调用SetTAG
	// 3. 需要建立连接
	for _,v := range tests {
		Factory.BroadCast(v.content,v.target...)
	}
	// v1  websocket output : hello content ,and v2,v3 no content output
	// v1,v2 websocket output : hello content,and v3 no content output
	// v1,v2,v3 websocket output : hello content

}



// 主要应对数据发送的时候版本的问题，比如某一条数据由于协议更改需要向上兼容老的版本,因为这是应用层的内容
// 所以使用wti 接口来进行兼容处理。避免进行内容的感染
func TestTg_BroadCastByTarget(t *testing.T) {
	tests := []struct {
		give map[string][]byte
	}{
		{
			give: map[string][]byte{
				"v1": []byte("first v1 "),
				"v2": []byte("second v2 "),
				"v3": []byte("third v3 "),
			},

		},
		{
			give: map[string][]byte{
				"v1": []byte("hello v1 "),
				"v2": []byte("hello v2 "),
				"v3": []byte("hello v3 "),
			},
		},

	}
	// 测试 v1、v2 、v1 三个版本的不同，如果要模式真实连接，则需要执行
	// 1. 创建im 连接服务器，并开启wti 配置参数
	// 2. 要在handler方法中调用factory 进行调用SetTAG
	// 3. 需要建立连接
	for _,v := range tests {
		Factory.BroadCastByTarget(v.give)
	}
	// v1,v2,v3  websocket output :first v1 | second v2 | third v3
	// v1,v2,v3  websocket output :hello v1 | hello v2 | hello v3
}




// 主要应对数据发送的时候版本的问题，比如某一条数据由于协议更改需要向上兼容老的版本,因为这是应用层的内容
// 所以使用wti 接口来进行兼容处理。避免进行内容的感染
func TestTg_UpdateAndF(t *testing.T) {
	tests := []struct {
		give string
	}{
		{
			give: "1234",
		},
	}
	// 这个测试条件相对比较苛刻，update接口主要作用是接收globalclosed 的信号，如果某个用户关闭连接
	// im线程就会释放连接之前就会告诉当前的update方法，所以只需要判断当前用户是否被删除就好了
	// 1. 创建im 连接服务器，并开启wti 配置参数
	// 2. 要在handler方法中调用factory 进行调用SetTAG
	// 3. 需要建立连接
	for _,v := range tests {
		Factory.Update(v.give) //
		fmt.Println(Factory.GetClienterTAGs(v.give))
	}
	// output : []
}