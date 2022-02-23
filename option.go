package im

import (
	"github.com/mongofs/im/client"
	"github.com/mongofs/im/log"
	"github.com/mongofs/im/plugins/validate"
	"github.com/mongofs/im/plugins/validate/example"
	"github.com/mongofs/im/plugins/wti"
)

const (
	// 对客户端进行默认参数设置
	DefaultClientHeartBeatInterval = 120
	DefaultClientReaderBufferSize  = 1024
	DefaultClientWriteBufferSize   = 1024
	DefaultClientBufferSize        = 8
	DefaultClientMessageType       = 1
	DefaultClientProtocol          = 1

	// 对分片进行基础设置
	DefaultBucketSize = 1 << 8 // 256

	// 默认基础的server配置
	DefaultServerBucketNumber = 1 << 6 // 64
	DefaultServerRpcPort      = ":8081"
	DefaultServerHttpPort     = ":8080"

	// 设置对广播能力的参数支持
	DefaultBroadCastHandler = 10
	DefaultBroadCastBuffer  = 200

	// plugins 的参数支持
	PluginWTISupport = false // 是否支持WTI 进行扩展
)

var DefaultValidate validate.Validater = &example.DefaultValidate{}
var DefaultReceive client.Receiver = &client.Example{}
var DefaultLogger log.Logger = &log.DefaultLog{}

type Option struct {
	// client
	ClientHeartBeatInterval int // 用户心跳间隔
	ClientReaderBufferSize  int // 用户连接读取buffer
	ClientWriteBufferSize   int // 用户连接写入buffer
	ClientBufferSize        int // 用户应用层buffer
	ClientMessageType       int // 用户发送的数据类型
	ClientProtocol          int // 压缩协议

	// bucket
	BucketSize         int // bucket用户

	// server
	ServerBucketNumber int // 所有
	ServerRpcPort      string
	ServerHttpPort     string
	ServerValidate     validate.Validater
	ServerReceive      client.Receiver
	ServerLogger       log.Logger

	//broadcast
	BroadCastBuffer  int
	BroadCastHandler int

	//plugins
	SupportPluginWTI bool // 是否支持wti插件
}

func DefaultOption() *Option {
	return &Option{
		ClientHeartBeatInterval: DefaultClientHeartBeatInterval,
		ClientReaderBufferSize:  DefaultClientReaderBufferSize,
		ClientWriteBufferSize:   DefaultClientWriteBufferSize,
		ClientBufferSize:        DefaultClientBufferSize,
		ClientMessageType:       DefaultClientMessageType,
		ClientProtocol:          DefaultClientProtocol,
		BucketSize:              DefaultBucketSize,

		ServerBucketNumber: DefaultServerBucketNumber, // 所有
		ServerRpcPort:      DefaultServerRpcPort,
		ServerHttpPort:     DefaultServerHttpPort,
		ServerValidate:     DefaultValidate,
		ServerReceive:      DefaultReceive,
		ServerLogger:       DefaultLogger,

		BroadCastBuffer:  DefaultBroadCastBuffer,
		BroadCastHandler: DefaultBroadCastHandler,

		// 插件支持
		SupportPluginWTI: PluginWTISupport,
	}
}

func NewOption(Opt ...OptionFunc) *Option {
	opt := DefaultOption()
	for _, o := range Opt {
		o(opt)
	}
	return opt
}

type OptionFunc func(b *Option)

func WithServerHttpPort(ServerHttpPort string) OptionFunc {
	return func(b *Option) {
		b.ServerHttpPort = ServerHttpPort
	}
}

func WithServerRpcPort(ServerRpcPort string) OptionFunc {
	return func(b *Option) {
		b.ServerRpcPort = ServerRpcPort
	}
}

func WithServerValidate(ServerValidate validate.Validater) OptionFunc {
	return func(b *Option) {
		b.ServerValidate = ServerValidate
	}
}

func WithServerLogger(ServerLogger log.Logger) OptionFunc {
	return func(b *Option) {
		b.ServerLogger = ServerLogger
	}
}

func WithServerBucketNumber(ServerBucketNumber int) OptionFunc {
	return func(b *Option) {
		b.ServerBucketNumber = ServerBucketNumber
	}
}

func WithServerReceive(ServerReceive client.Receiver) OptionFunc {
	return func(b *Option) {
		b.ServerReceive = ServerReceive
	}
}

func WithClientHeartBeatInterval(ClientHeartBeatInterval int) OptionFunc {
	return func(b *Option) {
		b.ClientHeartBeatInterval = ClientHeartBeatInterval
	}
}

func WithClientReaderBufferSize(ClientReaderBufferSize int) OptionFunc {
	return func(b *Option) {
		b.ClientReaderBufferSize = ClientReaderBufferSize
	}
}

func WithClientWriteBufferSize(ClientWriteBufferSize int) OptionFunc {
	return func(b *Option) {
		b.ClientWriteBufferSize = ClientWriteBufferSize
	}
}

func WithClientBufferSize(ClientBufferSize int) OptionFunc {
	return func(b *Option) {
		b.ClientBufferSize = ClientBufferSize
	}
}

func WithClientMessageType(ClientMessageType int) OptionFunc {
	return func(b *Option) {
		b.ClientMessageType = ClientMessageType
	}
}

func WithClientProtocol(ClientProtocol int) OptionFunc {
	return func(b *Option) {
		b.ClientProtocol = ClientProtocol
	}
}

func WithBucketSize(BucketSize int) OptionFunc {
	return func(b *Option) {
		b.BucketSize = BucketSize
	}
}

func WithBroadCastBuffer(BroadCastBuffer int) OptionFunc {
	return func(b *Option) {
		b.BroadCastBuffer = BroadCastBuffer
	}
}

func WithBroadCastHandler(BroadCastHandler int) OptionFunc {
	return func(b *Option) {
		b.BroadCastHandler = BroadCastHandler
	}
}

//设置plugin内容
func WithPluginsWTI(SupportPluginWTI bool) OptionFunc {
	return func(b *Option) {
		if SupportPluginWTI {
			wti.SetSupport()
		}
		b.SupportPluginWTI = SupportPluginWTI
	}
}
