package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	common "test/chatroom/common/message"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		// fmt.Println(string(buf[:n]))
		msg, err := common.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,服务端也退出")
				return
			}
			fmt.Println("readPkg err=", err)
			return
		}
		fmt.Println(msg)
		err = serverProcess(conn, &msg)
		if err != nil {
			fmt.Println("serverProcess err=", err)
			return
		}
	}
}

func serverLogin(msg *common.Message) (result common.LoginRes, err error) {
	// 取出msg Data
	var loginData common.LoginMes
	err = json.Unmarshal([]byte(msg.Data), &loginData)
	if err != nil {
		fmt.Println("反序列失败")
		return
	}
	// 判断用户id
	if loginData.UserId == 100 && loginData.UserPwd == "123456" {
		result.Code = 200
		result.Message = "登录成功"
	} else {
		result.Code = 500
		result.Message = "用户名或者密码错误"
	}

	return
}

// 根据客户端发送的消息类型分发函数
func serverProcess(conn net.Conn, msg *common.Message) (err error) {
	var resultMsg common.Message

	switch msg.Type {
	case common.LoginResponse:
		// 处理登录
		result, err := serverLogin(msg)
		if err != nil {
			return err
		}
		data, _ := json.Marshal(result)
		resultMsg.Type = common.LoginResult
		resultMsg.Data = string(data)
	case common.RegisterResponse:
	default:
		fmt.Println("找不到的类型", msg.Type)
	}

	fmt.Println(resultMsg)
	// 发送消息给客户端
	err = common.SendMessage(conn, resultMsg)

	return
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
