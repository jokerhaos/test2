package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	port := 8889

	listen, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))

	if err != nil {
		fmt.Println("服务启动失败,端口:", port)
		return
	}

	defer listen.Close()

	fmt.Println("服务器启动端口：0.0.0.0:", port)

	for {
		// 等待客户端的连接
		fmt.Println("等待客户端的连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("客户端的连接异常")
			continue
		}
		fmt.Println("客户端连接成功,客户端IP:", conn.RemoteAddr().String())
		// 启动一个协程处理此次连接
		go func(conn net.Conn) {
			defer conn.Close()
			for {
				fmt.Println("等待", conn.RemoteAddr().String(), "的消息")
				buf := make([]byte, 1024)
				// 如果对方没有conn.wrtie发消息会堵塞
				n, err := conn.Read(buf) // 从conn读取消息
				if err != nil {
					fmt.Println("conn.Read err=", err)
					return
				}
				fmt.Println(string(buf[:n]))
			}
		}(conn)
	}

	// fmt.Println("listen %v", listen)

}
