package client

import (
	"github.com/Rehtt/RehttKit/buf"
	"github.com/fatih/color"
)

type ShowOnline struct{}

func (s *ShowOnline) Key() string {
	return "showOnline"
}

func (s *ShowOnline) Help() string {
	return "显示其他在线的用户"
}

func (s *ShowOnline) Run(ctx *Context, src string) {
	for _, cc := range Group.OnlineUsers() {
		if cc == ctx {
			continue
		}
		ctx.Term.Write(buf.NewBuf().WriteString("$[").WriteString(cc.User.Name).WriteString("] ").
			WriteString("online\n").ToColorBytes([]color.Attribute{color.FgHiCyan}))
	}
}

func init() {
	RegisterCmd(&ShowOnline{})
}
