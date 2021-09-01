package terminal

import (
	"fmt"
	"gosh/utils"
)

func Run(ssh string) {
	//sercet := string(utils.AesGenerateKey(utils.AES256))
	utils.Cmd("ssh "+ssh, func(out string) (exit bool) {
		fmt.Print(out)
		return false
	})
}
