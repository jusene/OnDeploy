package models

// 请求
type ServerDetail struct {
	Name string `json:"name"`
	Address string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port uint `json:"port"`
	// Comment string `json:"comment,omitempty"`
}

type ServersDetail []ServerDetail

// 响应
type Res struct {
	Code uint `json:"code"`
	Message string `json:"msg"`
}

type Err Res