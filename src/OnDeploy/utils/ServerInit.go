package utils

import (
	"OnDeploy/models"
	"OnDeploy/templates"
	"errors"
	"fmt"
	"net"
	"time"
)

func InitServer(server models.ServerDetail, user, pass string) error {
	// 检查请求IP是否为IP格式
	if addr := net.ParseIP(server.Address); addr == nil {
		return errors.New(fmt.Sprintf("无效的IP地址: %s", server.Address))
	}

	// 检查主机是否存活
	conn, err := net.DialTimeout("ip:icmp", server.Address, 1*time.Second)
	if err != nil {
		return errors.New(fmt.Sprintf("服务器检测存活失败: %s", server.Address))
	}
	defer conn.Close()

	var client SSHClient
	client = NewClient(server, user, pass)
	defer client.Close()

	// 安装常用的应用包
	if _, err = client.RemoteExec("yum install -y wget vim ntpdate curl epel-release telnet git"); err != nil {
		return errors.New(fmt.Sprintf("%s 安装程序包错误: %v", server.Address, err))
	}

	// 设置服务器主机名
	if _, err = client.RemoteExec(fmt.Sprintf("hostnamectl set-hostname %s", server.Name)); err != nil {
		return errors.New(fmt.Sprintf("%s 设置服务器主机名错误: %v", server.Address, err))
	}

	// 设置服务器同步时间
	if _, err = client.RemoteExec(fmt.Sprintf("echo '*/5 * * * * root ntpdate ntp1.aliyun.com &> /dev/null' > /etc/cron.d/ntpdate ")); err != nil {
		return errors.New(fmt.Sprintf("%s 设置同步时间错误: %v", server.Address, err))
	}

	// 设置内核参数
	if _, err = client.RemotePut(templates.SysConfig, "/etc/sysctl.conf"); err != nil {
		return errors.New(fmt.Sprintf("%s 设置内核参数错误: %v", server.Address, err))
	}

	// 关闭防火墙
	if _, err = client.RemoteExec("systemctl disable firewalld && systemctl stop firewalld"); err != nil {
		return errors.New(fmt.Sprintf("%s 关闭防火墙失败: %v", server.Address, err))
	}

	// disable selinux
	if _, err = client.RemoteExec("setenforce 0 && sed -i 's/^SELINUX=.*/SELINUX=disabled/g' /etc/selinux/config"); err != nil {
		return errors.New(fmt.Sprintf("%s 关闭selinux失败: %v", server.Address, err))
	}

	// 关闭swap
	if _, err = client.RemoteExec("swapoff -a && sed -i 's/.*swap.*/#&/' /etc/fstab"); err != nil {
		return errors.New(fmt.Sprintf("%s 关闭swap失败: %v", server.Address, err))
	}

	return nil
}

func InitServers(server models.ServerDetail, user, pass string, errChan chan string) {
	err := InitServer(server, user, pass)
	if err != nil {
		errChan <- err.Error()
	}
}