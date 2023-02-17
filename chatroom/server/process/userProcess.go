package process

import (
	"encoding/json"
	"fmt"
	common "test/chatroom/common/message"
)

type UserProcess struct {
	Msg *common.Message
}

func (this *UserProcess) ServerLogin() (result common.LoginRes, err error) {
	// 取出msg Data
	var loginData common.LoginMes
	err = json.Unmarshal([]byte(this.Msg.Data), &loginData)
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
