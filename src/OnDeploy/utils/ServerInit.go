package utils

import (
	"OnDeploy/models"
	"OnDeploy/templates"
	"errors"
	"fmt"
	"net"
	"time"
)

func InitServer(server models.ServerDetail) error {
	// 检查请求IP是否为IP格式
	if addr := net.ParseIP(server.Address); addr == nil {
		return errors.New(fmt.Sprintf("无效的IP地址 %s", server.Address))
	}

	// 检查主机是否存活
	conn, err := net.DialTimeout("ip:icmp", server.Address, 1*time.Second)
	if err != nil {
		return errors.New(fmt.Sprintf("服务器检测存活失败 %s", server.Address))
	}
	defer conn.Close()

	var client SSHClient
	client = NewClient(server)
	// 安装常用的应用包
	if _, err = client.RemoteExec("yum install -y wget vim ntpdate sysstat curl epel-release telnet git"); err != nil {
		return errors.New(fmt.Sprintf("安装程序包错误: %v", err))
	}

	// 设置服务器主机名
	if _, err = client.RemoteExec(fmt.Sprintf("hostnamectl set-hostname %s", server.Name)); err != nil {
		return errors.New(fmt.Sprintf("设置服务器主机名错误: %v", err))
	}

	// 设置服务器同步时间
	if _, err = client.RemoteExec(fmt.Sprintf("echo '*/5 * * * * ntpdate ntp1.aliyun.com &> /dev/null' > ")); err != nil {
		return errors.New(fmt.Sprintf("设置同步时间错误: %v", err))
	}

	// 设置内核参数
	if _, err = client.RemotePut(templates.SysConfig, "/etc/sysctl.conf"); err != nil {
		return errors.New(fmt.Sprintf("设置内核参数错误: %v", err))
	}




	//

	// 关闭会话
	client.Close()
	return nil
}
