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

type NewRabbitVhost struct {
	Address string `json:"address"`
	Vhost string `json:"vhost"`
	Trace bool `json:"tracing"`
}

// permission
type RabbitPermission struct {
	User string `json:"user"`
	Vhost string `json:"vhost"`
	Configure string `json:"configure"`
	Write string `json:"write"`
	Read string `json:"read"`
}

type RabbitPermissions []RabbitPermission

type NewRabbitPermission struct {
	Address string `json:"address"`
	User string `json:"user"`
	Vhost string `json:"vhost"`
	Configure string `json:"configure"`
	Write string `json:"write"`
	Read string `json:"read"`
}

type ComRabbitPermission struct {
	Address string `json:"address"`
	User string `json:"user"`
	Vhost string `json:"vhost"`
}