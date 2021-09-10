## 使用示例

```go
package main

import (
	"fmt"

	customizeSSHClient "github.com/customize-ssh-client"
)

func main() {
	// 定义需要连接的远程主机信息（修改 IP 和端口信息）
	sshHost := "10.100.0.21"
	sshPort := "22"
	sshUser := "root"
	sshAddr := fmt.Sprintf("%v:%v", sshHost, sshPort)

	// 使用秘钥方式连接（修改私钥地址）
	id_rsa := "/root/.ssh/id_rsa"
	client := customizeSSHClient.SSHPrivateKey(sshAddr, sshUser, id_rsa)

	// 定义需要在远程服务器上执行的命令
	shellList := []string{
		"whoami",
		"ls -lha /root",
	}
	customizeSSHClient.ExecShell(client, shellList)
}
```

