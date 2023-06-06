package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"user/thrift/user"
)

func main() {
	// 打开连接
	socket := thrift.NewTSocketConf("localhost:9090", &thrift.TConfiguration{})
	if err := socket.Open(); err != nil {
		panic("server connect failed")
	}

	// 构造传输协议，要和服务端保持一致
	protocolFactory := thrift.NewTCompactProtocolFactoryConf(&thrift.TConfiguration{})
	// 创建客户端
	client := user.NewUserServiceClientFactory(socket, protocolFactory)
	// 发起调用请求
	ctx := context.Background()
	name, _ := client.GetName(ctx, 123)
	fmt.Println(name)

	name, err := client.GetName(ctx, -123)
	fmt.Println(err)

	// 关闭连接
	if err := socket.Close(); err != nil {
		panic("client close failed")
	}
}
