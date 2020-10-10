package models

// 请求
type ServerDetail struct {
	Name string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	//Username string `json:"username,omitempty"`
	//Password string `json:"password,omitempty"`
	Port uint `json:"port,omitempty"`
	// Comment string `json:"comment,omitempty"`
}

type ServersDetail []ServerDetail

// 响应
type Res struct {
	Code uint `json:"code"`
	Message string `json:"msg"`
}

type Err Res