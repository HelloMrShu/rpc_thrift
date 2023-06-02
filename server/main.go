package main

import (
	"github.com/apache/thrift/lib/go/thrift"
	"user/thrift/handler"
	"user/thrift/user"
)

func main() {
	ch := make(chan error)

	go func() {
		as := handler.UserService{}

		// 创建处理器
		processor := user.NewUserServiceProcessor(as)
		// 配置传输类型
		transportFactory := thrift.NewTBufferedTransportFactory(10000000)
		// 配置传输协议，序列化和反序列化，compact/binary/json
		protocolFactory := thrift.NewTCompactProtocolFactoryConf(&thrift.TConfiguration{})
		// 构造socket连接
		serverTransport, _ := thrift.NewTServerSocket("127.0.0.1:9090")
		// 配置服务器
		server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
		// 对外提供服务
		ch <- server.Serve()
	}()

	<-ch
}
