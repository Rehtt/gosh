package client

import (
	"fmt"
	"github.com/Rehtt/RehttKit/buf"
	"github.com/Rehtt/gosh/database"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
	"log"
	"strings"
)

type Context struct {
	Term         *term.Terminal
	User         *database.UserTable
	windowWidth  int
	windowHeight int
}

func NewClient(user *database.UserTable) *Context {
	return &Context{
		User: user,
	}
}

func (c *Context) Resize(width, height int) error {
	err := c.Term.SetSize(width, height)
	if err != nil {
		log.Printf("Resize failed: %dx%d", width, height)
		return err
	}
	c.windowWidth, c.windowHeight = width, height
	return nil
}

func (c *Context) HandleShell(channel ssh.Channel) {
	defer channel.Close()
	c.Term.Write([]byte("\033c"))
	c.Term.Write(buf.NewBuf().WriteString("Welcome ").WriteString(c.User.Name).WriteString(" !\n").ToBytes(true))
	c.Term.Write(buf.NewBuf().WriteString(strings.Join(Group.history.Range(), "")).ToBytes(true))

	com, _ := ParseCmd("/showOnline")
	com.Run(c, "")

	for {
		line, err := c.Term.ReadLine()
		if err != nil {
			break
		}

		fmt.Println(c.User.Name, line)
		if cm, isCmd := ParseCmd(line); isCmd {
			cm.Run(c, line)
		} else {
			Group.SendMsg(line, c)
		}
	}
}

func inArray(c *Context, cs []*Context) bool {
	for _, v := range cs {
		if c == v {
			return true
		}
	}
	return false
}
