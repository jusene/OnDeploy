package models

// user
type RabbitUser struct {
	Name string `json:"name"`
	Passwd string `json:"password_hash"`
	Tags string `json:"tags"`
}

type RabbitUsers []RabbitUser

type NewRabbitUser struct {
	Address string `json:"address"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

type CommonRabbitUser struct {
	Address string `json:"address"`
	User string `json:"user"`
}

// vhost
type RabbitVhost struct {
	Name string `json:"name"`
	Trace bool `json:"tracing"`
}

type RabbitVhosts []RabbitVhost