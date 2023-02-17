package common

const (
	LoginResponse    = "LoginResponse"
	LoginResult      = "LoginResult"
	RegisterResponse = "RegisterResponse"
	RegisterResult   = "RegisterResult"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginRes struct {
	Code    int
	Message string
}
