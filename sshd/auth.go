package sshd

import (
	"fmt"
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
			out, err := client(conn.User(), "zc:",
				[]string{"绑定：", "密码（输入不可见）:", "确认密码（输入不可见）:"},
				[]bool{true, false, false})
			if err != nil {
				return nil, err
			}
			if out[1] != out[2] {
				return nil, fmt.Errorf("no")
			}
			user.Name = conn.User()
			user.Password = &out[1]
			user.QNo = out[0]
		} else {
			out, err := client(conn.User(), "login:",
				[]string{"密码（输入不可见）:"},
				[]bool{false})
			if err != nil {
				return nil, err
			}
			if out[0] != *user.Password {
				return nil, fmt.Errorf("mmcw")
			}
		}
		user.LastLoginTime = time.Now()
		user.Save(db)

		return &ssh.Permissions{
			Extensions: map[string]string{"userJson": user.ToJson()},
		}, nil
	}
}
