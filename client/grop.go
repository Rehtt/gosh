package client

import (
	"github.com/Rehtt/RehttKit/buf"
	"github.com/Rehtt/gosh/util"
	"github.com/fatih/color"
	"sync"
)

type GroupStruct struct {
	users   map[string]*Context
	history *util.Link
}

var (
	Group   *GroupStruct
	GroupMu sync.RWMutex
)

func init() {
	NewGroup()
}

func NewGroup() {
	h := util.NewLink()
	h.SetMaxSize(10)
	Group = &GroupStruct{
		users:   make(map[string]*Context),
		history: h,
	}
}

func (g *GroupStruct) Login(ctx *Context) {
	GroupMu.Lock()
	if g.users == nil {
		g.users = make(map[string]*Context)
	}
	g.users[ctx.User.Name] = ctx
	GroupMu.Unlock()

	g.SendOnline(ctx)
}
func (g *GroupStruct) Logout(ctx *Context) {
	GroupMu.Lock()
	delete(g.users, ctx.User.Name)
	GroupMu.Unlock()

	g.SendOffline(ctx)
}
func (g *GroupStruct) OnlineUsers() (c []*Context) {
	GroupMu.RLock()
	defer GroupMu.RUnlock()
	c = make([]*Context, 0, len(g.users))
	for k := range g.users {
		c = append(c, g.users[k])
	}
	return
}
func (g *GroupStruct) Send(data []byte, ignores ...*Context) {
	for _, cc := range Group.OnlineUsers() {
		if cc.Term == nil || inArray(cc, ignores) {
			continue
		}
		cc.Term.Write(data)
	}
}
func (g *GroupStruct) SendMsg(msg string, from *Context) {
	b := buf.NewBuf().WriteString("[").
		WriteString(from.User.Name).
		WriteString("]: ").
		WriteString(msg).
		WriteByte('\n')
	g.history.Write(b.ToString())
	g.Send(b.ToBytes(), from)
}
func (g *GroupStruct) SendOnline(ctx *Context) {
	b := buf.NewBuf().WriteString("DING:").
		WriteString(" [").WriteString(ctx.User.Name).WriteString("] online\n")

	g.Send(b.ToColorBytes([]color.Attribute{color.FgHiGreen}, true))
}
func (g *GroupStruct) SendOffline(ctx *Context) {
	b := buf.NewBuf().WriteString("DONG:").
		WriteString(" [").WriteString(ctx.User.Name).WriteString("] offline\n")

	g.Send(b.ToColorBytes([]color.Attribute{color.FgHiRed}, true))
}
