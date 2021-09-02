package terminal

import (
	"fmt"
	"gosh/utils"
	"strings"
)

func Run(cmd string) {
	sercet := string(utils.AesGenerateKey(utils.AES256))
	utils.SSH(fmt.Sprintf(`%s "gosh -secret %s&"`, cmd, sercet), func(out string) {
		if strings.Contains(out, "gosh: command not found") {
			fmt.Println("not find gosh")
			return
		}
		fmt.Println(out)
	})
}
