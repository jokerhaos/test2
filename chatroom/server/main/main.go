package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		fmt.Println("服务启动失败", err)
		return
	}
	fmt.Println("服务器启动成功 0.0.0.0:9000")
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
			continue
		}
		// 连接成功启动协程和客户端交互
		process := &Processor{
			Conn: conn,
		}
		go process.ProcessMsg()
	}

}
