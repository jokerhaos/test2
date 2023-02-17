package main

import (
	"encoding/json"
	"fmt"
	"net"
	common "test/chatroom/common/message"
	process "test/chatroom/server/process"
	utils "test/chatroom/util"
)

// 根据客户端发送的消息类型分发函数
func ServerProcess(conn net.Conn, msg *common.Message) (err error) {
	var resultMsg common.Message

	switch msg.Type {
	case common.LoginResponse:
		// 处理登录
		result, err := process.ServerLogin(msg)
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
	err = utils.SendMessage(conn, resultMsg)

	return
}
