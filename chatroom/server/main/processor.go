package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	common "test/chatroom/common/message"
	process "test/chatroom/server/process"
	utils "test/chatroom/util"
)

type Processor struct {
	Conn net.Conn
}

// 根据客户端发送的消息类型分发函数
func (this *Processor) ServerProcess(msg *common.Message) (err error) {
	var resultMsg common.Message

	switch msg.Type {
	case common.LoginResponse:
		// 处理登录
		userProcess := &process.UserProcess{
			Msg: msg,
		}
		result, err := userProcess.ServerLogin()
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
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.SendMessage(resultMsg)

	return
}

func (this *Processor) ProcessMsg() {
	defer this.Conn.Close()
	for {
		// fmt.Println(string(buf[:n]))
		tf := &utils.Transfer{
			Conn: this.Conn,
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
		err = this.ServerProcess(&msg)
		if err != nil {
			fmt.Println("serverProcess err=", err)
			return
		}
	}
}
