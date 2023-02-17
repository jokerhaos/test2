package main

import (
	"fmt"
	"io"
	"net"
	utils "test/chatroom/util"
)

func processMsg(conn net.Conn) {
	defer conn.Close()
	for {
		// fmt.Println(string(buf[:n]))
		tf := &utils.Transfer{
			Conn: conn,
		}
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,服务端也退出")
				return
			}
			fmt.Println("readPkg err=", err)
			return
		}
		fmt.Println(msg)
		err = ServerProcess(conn, &msg)
		if err != nil {
			fmt.Println("serverProcess err=", err)
			return
		}
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
		go processMsg(conn)
	}

}
