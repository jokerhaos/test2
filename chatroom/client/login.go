package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	common "test/chatroom/common/message"
)

// 登录
func login(userId int, userPwd string) (err error) {
	// 定义协议
	fmt.Printf("userId=%d userPwd=%s", userId, userPwd)

	// 和服务端建立通讯
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("和服务端建立通讯异常", err)
		return err
	}

	defer conn.Close()

	// 登录参数
	var data common.LoginMes
	data.UserId = userId
	data.UserPwd = userPwd
	jsonResult, _ := json.Marshal(data)

	// 消息参数
	var msg common.Message
	msg.Type = common.LoginResponse
	msg.Data = string(jsonResult)

	// 给服务端发送消息
	err = common.SendMessage(conn, msg)
	if err != nil {
		fmt.Println("send err=", err)
		return err
	}

	// 接收服务端消息
	msg, err = common.ReadPkg(conn)
	if err != nil {
		fmt.Println("receive err=", err)
		return
	}
	var result common.LoginRes
	json.Unmarshal([]byte(msg.Data), &result)
	if result.Code != 200 {
		return errors.New(result.Message)
	}
	return
}
