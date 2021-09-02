package terminal

import (
	"fmt"
	"gosh/utils"
	"strings"
)

var (
	port   string
	secret = string(utils.AesGenerateKey(utils.AES256))
)

func Run(cmd string) {
	utils.SSH(fmt.Sprintf(`%s "gosh -secret %s &"`, cmd, secret), func(out string) (exit bool) {
		if strings.Contains(out, "gosh: command not found") {
			fmt.Println("not find gosh")
			return
		}
		if strings.Contains(out, "success") {
			port = strings.Split(out, " ")[1]
			return true
		}
		if strings.Contains(out, "error") {
			fmt.Println("failed to open port")
			return true
		}
		fmt.Println(out)
		return false
	})
	fmt.Println(port)
}
