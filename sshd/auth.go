package sshd

import (
	"github.com/Rehtt/gosh/database"
	"golang.org/x/crypto/ssh"
	"gorm.io/gorm"
	"time"
)

func authKeyboard(db *gorm.DB) func(conn ssh.ConnMetadata, client ssh.KeyboardInteractiveChallenge) (*ssh.Permissions, error) {
	return func(conn ssh.ConnMetadata, client ssh.KeyboardInteractiveChallenge) (*ssh.Permissions, error) {
		username := conn.User()
		user := (&database.UserTable{}).Get(db, username)
		if user.ID == 0 {
			out, err := client(conn.User(), "Register",
				[]string{"绑定QQ："},
				[]bool{true})
			if err != nil {
				return nil, err
			}
			user.QNo = out[0]
			user.Name = conn.User()
			instruction := ""
			for {

				out, err = client(conn.User(), instruction, []string{"密码（输入不可见）:", "确认密码（输入不可见）:"}, []bool{false, false})
				if err != nil {
					return nil, err
				}
				if out[0] != out[1] {
					instruction = "Inconsistent password"
				} else {
					user.Password = &out[1]
					break
				}
			}

		} else {
			instruction := "Login"
			for {
				out, err := client(conn.User(), instruction,
					[]string{"密码（输入不可见）:"},
					[]bool{false})
				if err != nil {
					return nil, err
				}
				if out[0] != *user.Password {
					instruction = "Cross -password error, try again"
				} else {
					break
				}
			}
		}
		user.LastLoginTime = time.Now()
		user.Save(db)

		return &ssh.Permissions{
			Extensions: map[string]string{"userJson": user.ToJson()},
		}, nil
	}
}
