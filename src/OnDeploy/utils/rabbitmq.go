package utils

import (
	"OnDeploy/models"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func InstallRabbitMQ(server models.ServerDetail) error {
	var client SSHClient
	client = NewClient(server)
	defer client.Close()

	if _, err := client.RemoteExec("yum install -y rabbitmq-server"); err != nil {
		return errors.New(fmt.Sprintf("%s 安装rabbitmq程序包错误: %v", server.Address, err))
	}

	if _, err := client.RemoteExec("systemctl enable rabbitmq-server --now"); err != nil {
		return errors.New(fmt.Sprintf("%s 启动rabbitmq错误: %v", server.Address, err))
	}

	if _, err := client.RemoteExec("rabbitmq-plugins enable rabbitmq_management && systemctl restart rabbitmq-server"); err != nil {
		return errors.New(fmt.Sprintf("%s 启动rabbitmq_management错误: %v", server.Address, err))
	}

	return nil
}


type Rabbit interface {
	LstUser() (string, error)
	AddUser(username, password string) error
	DelUser(username string) error
	LstVhost() (string, error)
	AddVhost(vhost string) error
	DelVhost(vhost string) error
}

type RabbitApi struct {
	Url string
	User string
	Pass string
}

func NewRabbitApi(address, username, password string, port int) *RabbitApi {
	return &RabbitApi{Url: fmt.Sprintf("http://%s:%d", address, port),
		User: username,
		Pass: password}
}

func (r *RabbitApi) LstUser() (string, error){
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/users", r.Url), nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(r.User, r.Pass)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	return string(data), nil
}

func (r *RabbitApi) AddUser(username, password string) error {
	client := &http.Client{}
	body := bytes.NewReader([]byte(fmt.Sprintf(`{"password":"%s","tags":"administrator"}`, password)))
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/users/%s", r.Url, username), body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(r.User, r.Pass)
	resp, err := client.Do(req)
	if err != nil {
		return  err
	}
	defer resp.Body.Close()
	return  nil
}

func (r *RabbitApi) DelUser(username string) error {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/users/%s", r.Url, username), nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(r.User, r.Pass)
	resp, err := client.Do(req)
	if err != nil {
		return  err
	}
	defer resp.Body.Close()
	return nil
}

func (r *RabbitApi) LstVhost() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/vhosts", r.Url), nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(r.User, r.Pass)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	return string(data), nil
}

func (r *RabbitApi) AddVhost(vhost string) error {
	return nil
}

func (r *RabbitApi) DelVhost(vhost string) error {
	return nil
}