package models

type NFSPath struct {
	ServerDetail
	PATH string `json:"path"`
	ACL string `json:"acl"`
	Param []string `json:"param"`
}

type NFSInfo struct {
	Address string `json:"address"`
	PATH string `json:"path"`
	Port uint `json:"port"`
}