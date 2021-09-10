package sshClient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

// 使用密钥方式连接 SSH 服务器
func SSHPrivateKey(sshAddr string, sshUser string, id_rsa string) *ssh.Client {
	// Public Key 可用于使用未加密的 PEM 编码的私钥文件对远程服务器进行身份验证
	// 如果你有一个加密的秘钥，可以使用 crypto/x509 包来解密他
	key, err := ioutil.ReadFile(id_rsa)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// 为此私钥创建签名者
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			// 使用 PublicKeys 方法进行元整身份验证
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 连接到远程服务器并执行 SSH 握手
	client, err := ssh.Dial("tcp", sshAddr, config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}

	return client
}

// 使用密码方式连接 SSH 服务器
func SSHPassword(sshAddr string, sshUser string, password string) *ssh.Client {
	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			// 使用 PublicKeys 方法进行元整身份验证
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 连接到远程服务器并执行 SSH 握手
	client, err := ssh.Dial("tcp", sshAddr, config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}

	return client
}

// 执行定义的命令
func ExecShell(client *ssh.Client, shellList []string) {
	// 延迟关闭连接
	defer client.Close()

	// 循环执行命令
	for index, shell := range shellList {
		// 每个 ClientConn 可以支持多个交互会话，由一个 Session 表示
		session, err := client.NewSession()
		if err != nil {
			log.Fatal("Failed to create session: ", err)
		}
		defer session.Close()
		// 创建会话后，你可以使用 Run 方法在远程端执行单个命令
		var b bytes.Buffer
		session.Stdout = &b
		fmt.Println("------------------------- 执行第", index+1, "条命令：", shell)
		if err := session.Run(shell); err != nil {
			log.Fatal("Failed to run: " + err.Error())
		}
		// 打印命令结果
		fmt.Println(b.String())
	}
}
