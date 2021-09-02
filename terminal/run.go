package terminal

import (
	"fmt"
	"gosh/utils"
	"strings"
)

func Run(cmd string) {
	sercet := string(utils.AesGenerateKey(utils.AES256))
	ssh(fmt.Sprintf(`%s "gosh -secret %s&"`, cmd, sercet), func(out string) {
		if strings.Contains(out, "gosh: command not found") {
			fmt.Println("not find gosh")
			return
		}
		fmt.Println(out)
	})
}

func ssh(cmd string, o func(out string)) error {
	a, err := utils.Cmd("ssh "+cmd, func(out string) (exit bool) {
		o(out)
		return false
	})
	err = a.Wait()
	return err
}
