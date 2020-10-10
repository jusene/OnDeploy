package utils

import (
	"OnDeploy/models"
	"errors"
	"fmt"
	"strings"
)

func InstallNFS(server models.ServerDetail, user, pass string) error {
	var client SSHClient
	client = NewClient(server, user, pass)
	defer client.Close()

	// 安装nfs程序包
	if _, err := client.RemoteExec("yum install -y nfs-utils"); err != nil {
		return errors.New(fmt.Sprintf("%s 安装nfs程序包错误: %v", server.Address, err))
	}

	// 启动服务
	if ret, err := client.RemoteExec("systemctl enable nfs --now"); err != nil {
		if !strings.Contains(ret, "Created symlink") {
			return errors.New(fmt.Sprintf("%s 启动nfs错误: %v", server.Address, err))
		}
	}

	return nil
}

type NFS interface {
	AddPath(nfspath models.NFSPath) error
	LstPath() (string, error)
	DelPath(nfspath models.NFSPath) error
}

type NFSObj struct {
	Server models.NFSPath
	User string
	Pass string
}

func (n *NFSObj) AddPath(nfspath models.NFSPath) error {
	var client SSHClient
	client = NewClient(n.Server.ServerDetail, n.User, n.Pass)
	defer client.Close()

	if _, err := client.RemoteExec(fmt.Sprintf("echo '%s %s(%s)' >> /etc/exports && mkdir -p %s", nfspath.PATH, nfspath.ACL,
		strings.Join(nfspath.Param, ","), nfspath.PATH)); err != nil {
		return errors.New(fmt.Sprintf("写入配置%s时失败", nfspath.PATH))
	}

	if _, err := client.RemoteExec("exportfs -rv"); err != nil {
		return errors.New("重载配置时失败")
	}

	return nil
}

func (n *NFSObj) LstPath() (string, error) {
	var client SSHClient
	client = NewClient(n.Server.ServerDetail, n.User, n.Pass)
	defer client.Close()

	ret, err := client.RemoteExec("cat /etc/exports")
	if err != nil {
		return "", errors.New("查看配置时失败")
	}

	return ret, nil
}

func (n *NFSObj) DelPath(nfspath models.NFSPath) error {
	var client SSHClient
	client = NewClient(n.Server.ServerDetail, n.User, n.Pass)
	defer client.Close()

	if ret, err := client.RemoteExec(fmt.Sprintf("sed -i 's@^%s.*@@g' /etc/exports && rm -rf %s",
		nfspath.PATH, nfspath.PATH)); err != nil {
		fmt.Println(ret)
		return errors.New(fmt.Sprintf("删除%s时失败", nfspath.PATH))
	}

	if _, err := client.RemoteExec("exportfs -rv"); err != nil {
		return errors.New("重载配置时失败")
	}

	return nil
}