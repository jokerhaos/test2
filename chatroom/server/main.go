package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 8096)
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		fmt.Println(buf[:n])
		// fmt.Println(string(buf[:n]))

	}
}

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
		go process(conn)
	}

}
