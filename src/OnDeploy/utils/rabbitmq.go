package utils

import (
	"OnDeploy/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func InstallRabbitMQ(server models.ServerDetail, user, pass string) error {
	var client SSHClient
	client = NewClient(server, user, pass)
	defer client.Close()

	if _, err := client.RemoteExec("yum install -y rabbitmq-server"); err != nil {
		return errors.New(fmt.Sprintf("%s 安装rabbitmq程序包错误: %v", server.Address, err))
	}

	if ret, err := client.RemoteExec("systemctl enable rabbitmq-server --now"); err != nil {
		if !strings.Contains(ret, "Created symlink") {
			return errors.New(fmt.Sprintf("%s 启动rabbitmq错误: %v", server.Address, err))
		}
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
	AddVhost(vhost string, trace bool) error
	DelVhost(vhost string) error
	LstPermission() (string, error)
	AddPermission(permission *models.NewRabbitPermission) error
	DelPermission(permission *models.ComRabbitPermission) error
}

type RabbitApi struct {
	Url  string
	User string
	Pass string
	// http client object
	HttpClient *http.Client
}

func NewRabbitApi(address, username, password string, port int) *RabbitApi {
	return &RabbitApi{Url: fmt.Sprintf("http://%s:%d", address, port),
		User:       username,
		Pass:       password,
		HttpClient: &http.Client{}}
}

func (r *RabbitApi) LstUser() (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/users", r.Url), nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(r.User, r.Pass)
	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	return string(data), nil
}

func (r *RabbitApi) AddUser(username, password string) error {
	body := bytes.NewReader([]byte(fmt.Sprintf(`{"password":"%s","tags":"administrator"}`, password)))
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/users/%s", r.Url, username), body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(r.User, r.Pass)
	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (r *RabbitApi) DelUser(username string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/users/%s", r.Url, username), nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(r.User, r.Pass)
	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (r *RabbitApi) LstVhost() (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/vhosts", r.Url), nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(r.User, r.Pass)
	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	return string(data), nil
}

func (r *RabbitApi) AddVhost(vhost string, trace bool) error {
	body := bytes.NewReader([]byte(fmt.Sprintf(`{"tracing": "%v"}`, trace)))
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/vhosts/%s", r.Url, url.QueryEscape(vhost)), body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(r.User, r.Pass)
	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (r *RabbitApi) DelVhost(vhost string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/vhosts/%s", r.Url, url.QueryEscape(vhost)), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(r.User, r.Pass)
	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

func (r *RabbitApi) LstPermission() (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/permissions", r.Url), nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(r.User, r.Pass)
	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	return string(data), nil
}

func (r *RabbitApi) AddPermission(permission *models.NewRabbitPermission) error {
	// 检查是否有请求的vhost存在
	if vhosts, err := r.LstVhost(); err != nil {
		return err
	} else {
		var v models.RabbitVhosts
		err = json.Unmarshal([]byte(vhosts), &v)
		if err != nil {
			return err
		}

		for _, vhost := range v {
			if permission.Vhost == vhost.Name {
				goto NEXTCHECK
			}
		}
		return errors.New(fmt.Sprintf("%s vhost不存在", permission.Vhost))
	}
NEXTCHECK:

	// 检查是否有请求的的user存在
	if users, err := r.LstUser(); err != nil {
		return err
	} else {
		var u models.RabbitUsers
		err = json.Unmarshal([]byte(users), &u)
		if err != nil {
			return err
		}

		for _, user := range u {
			if permission.User == user.Name {
				goto NEXT
			}
		}
		return errors.New(fmt.Sprintf("%s user不存在", permission.User))
	}
NEXT:

	body := bytes.NewReader([]byte(fmt.Sprintf(`{"configure":"%s","write":"%s","read":"%s"}`,
		permission.Configure, permission.Write, permission.Read)))

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/permissions/%s/%s", r.Url,
		url.QueryEscape(permission.Vhost), permission.User), body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(r.User, r.Pass)
	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (r *RabbitApi) DelPermission(permission *models.ComRabbitPermission) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/permissions/%s/%s", r.Url,
		url.QueryEscape(permission.Vhost), permission.User), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(r.User, r.Pass)
	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}
