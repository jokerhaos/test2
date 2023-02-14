package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("连接服务器失败")
		return
	}

	reader := bufio.NewReader(os.Stdin) // os.Stdin 标准的终端输入

	for {
		// 从终端读取一行用户输入并准备发送到服务端
		line, _ := reader.ReadString('\n')

		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出...")
			break
		}

		_, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("发送消息失败了 err=", err)
		}

		// fmt.Printf("客户端发送了 %d 字节数据", n)
	}

}
