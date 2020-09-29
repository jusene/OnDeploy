package utils

import (
	"OnDeploy/models"
	"errors"
	"fmt"
	)

func InstallNFS(server models.ServerDetail) error {
	var client SSHClient
	client = NewClient(server)
	defer client.Close()

	// 安装nfs程序包
	if _, err := client.RemoteExec("yum install -y nfs-utils"); err != nil {
		return errors.New(fmt.Sprintf("%s 安装nfs程序包错误: %v", server.Address, err))
	}

	// 启动服务
	if _, err := client.RemoteExec("systemctl enable nfs --now"); err != nil {
		return errors.New(fmt.Sprintf("%s 启动nfs错误: %v", server.Address, err))
	}

	return nil
}
