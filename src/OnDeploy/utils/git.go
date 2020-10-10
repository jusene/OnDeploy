package utils

import (
	"OnDeploy/models"
	"errors"
	"fmt"
	"strings"
)

func InstallGit(gitdetail models.GITDetail, user, pass, gitConfig string) error {
	var client SSHClient
	client = NewClient(gitdetail.ServerDetail, user, pass)
	defer client.Close()

	// 安装程序包
	if _, err := client.RemoteExec("yum install -y httpd git"); err != nil {
		return errors.New(fmt.Sprintf("%s 安装git程序包错误: %v", gitdetail.Address, err))
	}

	// 配置程序包
	if _, err := client.RemotePut(gitConfig, "/etc/httpd/conf.d/git.conf"); err != nil {
		return errors.New(fmt.Sprintf("配置git错误: %v", err))
	}

	// 创建git的工作目录
	if _, err := client.RemoteExec(fmt.Sprintf(
		"mkdir -p /ddhome/local/gitdata && " +
			"git init --bare /ddhome/local/gitdata/%s.git && " +
			"chown -R apache.apache /ddhome/local/gitdata/ && " +
			"cd /ddhome/local/gitdata/%s.git;git config http.receivepack true", gitdetail.GITRepo, gitdetail.GITRepo)); err != nil {
		return errors.New(fmt.Sprintf("创建git工作目录错误: %v", err))
	}

	// 配置git用户名密码
	if ret, err := client.RemoteExec(fmt.Sprintf("htpasswd -b -c -m /etc/httpd/conf/.httpd %s %s", gitdetail.GITUser, gitdetail.GITPass)); err != nil {
		if !strings.Contains(ret, "Adding password for user") {
			return errors.New(fmt.Sprintf("创建git用户名密码错误: %v", err))
		}
	}

	// 启动git服务
	if ret, err := client.RemoteExec("systemctl enable httpd --now"); err != nil {
		if !strings.Contains(ret, "Created symlink") {
			return errors.New(fmt.Sprintf("启动git服务错误: %v", err))
		}
	}

	return nil
}


type GIT interface {
	AddRepo(repo string) error
	DelRepo(repo string) error
	AddUser(user, pass string) error
	DelUser(user string) error
	UpdateUser(user, pass string) error
}

type GITObj struct {
	Server models.GITDetail
	User string
	Pass string
}

func (g *GITObj) AddRepo(repo string) error {
	var client SSHClient
	client = NewClient(g.Server.ServerDetail, g.User, g.Pass)
	defer client.Close()

	if _, err := client.RemoteExec(fmt.Sprintf(
		"git init --bare /ddhome/local/gitdata/%s.git && " +
			"chown -R apache.apache /ddhome/local/gitdata/ && " +
			"cd /ddhome/local/gitdata/%s.git;git config http.receivepack true", repo, repo)); err != nil {
		return errors.New(fmt.Sprintf("创建git仓库错误: %v", err))
	}
	return nil
}

func (g *GITObj) DelRepo(repo string) error {
	var client SSHClient
	client = NewClient(g.Server.ServerDetail, g.User, g.Pass)
	defer client.Close()

	if _, err := client.RemoteExec(fmt.Sprintf("rm -rf /ddhome/local/gitdata/%s.git", repo)); err != nil {
		return errors.New(fmt.Sprintf("删除git仓库错误: %v", err))
	}

	return nil
}

func (g *GITObj) AddUser(user, pass string) error {
	var client SSHClient
	client = NewClient(g.Server.ServerDetail, g.User, g.Pass)
	defer client.Close()

	if _, err := client.RemoteExec(fmt.Sprintf("htpasswd -b -m /etc/httpd/conf/.httpd %s %s", user, pass)); err != nil {
		return errors.New(fmt.Sprintf("添加git仓库用户错误: %v", err))
	}
	return nil
}

func (g *GITObj) DelUser(user string) error {
	var client SSHClient
	client = NewClient(g.Server.ServerDetail, g.User, g.Pass)
	defer client.Close()

	if _, err := client.RemoteExec(fmt.Sprintf("sed -i '/^%s/d' /etc/httpd/conf/.httpd", user)); err != nil {
		return errors.New(fmt.Sprintf("删除git用户失败: %v", err))
	}
	return nil
}

func (g *GITObj) UpdateUser(user, pass string) error {
	var client SSHClient
	client = NewClient(g.Server.ServerDetail, g.User, g.Pass)
	defer client.Close()

	// 判断是否有此用户
	ret, err := client.RemoteExec("cat /etc/httpd/conf/.httpd")
	if err != nil {
		return errors.New(fmt.Sprintf("未找到密码配置文件: %v", err))
	}

	for _, u := range strings.Split(ret, "\n") {
		if len(u) == 1 {
			continue
		}
		if strings.Split(u, ":")[0] == user {
			goto NEXT
		}
	}
	return errors.New(fmt.Sprintf("用户不存在: %v", err))
	NEXT:

	if _, err := client.RemoteExec(fmt.Sprintf("htpasswd -b -m /etc/httpd/conf/.httpd %s %s", user, pass)); err != nil {
		return errors.New(fmt.Sprintf("git用户密码修改失败: %v", err))
	}
	return nil
}