package utils

import (
	"OnDeploy/models"
	"errors"
	"fmt"
	)

func InstallRabbitMQ(server models.ServerDetail) error {
	var client SSHClient
	client = NewClient(server)
	defer client.Close()

	if _, err := client.RemoteExec("yum install -y epel-release rabbitmq-server"); err != nil {
		return errors.New(fmt.Sprintf("%s 安装rabbitmq程序包错误: %v", server.Address, err))
	}

	if _, err := client.RemoteExec("systemctl enable rabbitmq-server --now"); err != nil {
		return errors.New(fmt.Sprintf("%s 启动rabbitmq错误: %v", server.Address, err))
	}

	return nil
}
