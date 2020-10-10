package utils

import (
	"OnDeploy/models"
	"bytes"
	"errors"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"time"
)

type SSHClient interface {
	RemoteExec(cmd string) (ret string, err error)
	RemotePutFile(src, dest string) (ret string, err error)
	RemoteGetFile(src, dest string) (ret string, err error)
	RemotePut(src, dest string) (ret string, err error)
	Close()
}

type Client struct {
	sshClient *ssh.Client
	sftpClient *sftp.Client
}

func NewClient(server models.ServerDetail, user, pass string) *Client {
	conf := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         3 * time.Second,
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", server.Address, server.Port), conf)
	if err != nil {
		panic(fmt.Sprintf("创建ssh连接 %s:%d 失败: %v", server.Address, server.Port, err))
	}

	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		panic(fmt.Sprintf("创建sftp连接 %s:%d 失败: %v", server.Address, server.Port, err))
	}

	return &Client{
		sshClient: sshClient,
		sftpClient: sftpClient,
	}
}

func (c Client) RemoteExec(cmd string) (ret string, err error) {
	session, err := c.sshClient.NewSession()
	if err != nil {
		return "", errors.New("创建ssh session失败")
	}
	defer session.Close()

	var stdOut, stdErr bytes.Buffer
	session.Stdout = &stdOut
	session.Stderr = &stdErr

	session.Run(cmd)
	if stdErr.String() != "" {
		return stdErr.String(), errors.New(cmd)
	} else {
		return stdOut.String(), nil
	}
}

func (c Client) RemotePutFile(src, dest string) (ret string, err error) {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return "", errors.New(fmt.Sprintf("打开源文件错误: %s, %v", src, err))
	}
	defer srcFile.Close()

	// 创建目的文件
	// 先清理目标的文件
	c.RemoteExec(fmt.Sprintf("rm -rf %s", dest))
	destFile, err := c.sftpClient.Create(dest)
	if err != nil {
		return "", errors.New(fmt.Sprintf("创建目标文件错误: %s, %v", src, err))
	}
	defer destFile.Close()

	// 读取源文件
	fileCopy, err := ioutil.ReadAll(srcFile)
	if err != nil {
		return "", errors.New(fmt.Sprintf("源文件读取错误: %s, %v", src, err))
	}

	// 写入目标文件
	destFile.Write(fileCopy)
	return fmt.Sprintf("%s 文件传输成功 => %s", src, dest), nil
}

func (c Client) RemoteGetFile(src, dest string) (ret string, err error) {
	// 源文件打开
	srcFile, err := c.sftpClient.Open(src)
	if err != nil {
		return "", errors.New(fmt.Sprintf("打开源文件错误: %s, %v", src, err))
	}
	defer srcFile.Close()

	// 创建目标文件
	// 先清除本地文件
	os.Remove(dest)
	destFile, err := os.Create(dest)
	if err != nil {
		return "", errors.New(fmt.Sprintf("创建目标文件错误: %s, %v", src, err))
	}
	defer destFile.Close()

	if _, err := srcFile.WriteTo(destFile); err != nil {
		return "", errors.New(fmt.Sprintf("文件传输错误: %s, %v", src, err))
	}

	return fmt.Sprintf("%s 文件传输成功 => %s", src, dest), nil

}

func (c Client) RemotePut(src, dest string) (ret string, err error) {
	// 创建目标文件
	// 先清除
	c.RemoteExec(fmt.Sprintf("rm -rf %s", dest))
	destFile, err := c.sftpClient.Create(dest)
	if err != nil {
		return "", errors.New(fmt.Sprintf("创建目标文件错误: %s, %v", src, err))
	}
	defer destFile.Close()

	destFile.Write([]byte(src))
	return fmt.Sprintf("文件传输成功 => %s", dest), nil
}

func (c Client) Close() {
	c.sftpClient.Close()
	c.sshClient.Close()
}

