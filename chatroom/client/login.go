package main

import (
	"encoding/binary"
	"encoding/json"
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
	sendData, _ := json.Marshal(msg)

	// 1.先发送长度，防止丢包
	var pkgLen uint32
	pkgLen = uint32(len(sendData))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	conn.Write(buf[:4])
	fmt.Println("发送长度:", pkgLen)

	// 2.发送消息数据
	conn.Write(sendData)

	return
}
