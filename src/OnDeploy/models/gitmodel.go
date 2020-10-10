package models

type GITDetail struct {
	ServerDetail
	GITPath string `json:"git_path"`
	GITRepo string `json:"git_repo"`
	GITPort uint `json:"git_port"`
	GITUser string `json:"git_user"`
	GITPass string `json:"git_pass"`
}


type GITRepo struct {
	Address string `json:"address"`
	Repo string `json:"git_repo"`
	Port uint `json:"port"`
}

type GITUser struct {
	Address string `json:"address"`
	User string `json:"user"`
	Port uint `json:"port"`
}

type GITUserPass struct {
	Address string `json:"address"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Port uint `json:"port"`
}