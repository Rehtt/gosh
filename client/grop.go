package client

import (
	"github.com/Rehtt/Kit/buf"
	"github.com/Rehtt/gosh/database"
	"github.com/fatih/color"
	"log"
	"sort"
	"sync"
)

type GroupStruct struct {
	users map[string]*Context
}

var (
	Group   *GroupStruct
	GroupMu sync.RWMutex
)

func init() {
	NewGroup()
}

func NewGroup() {
	Group = &GroupStruct{
		users: make(map[string]*Context),
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

	// todo use sync.pool
	(&database.GroupTable{
		UserID:   from.User.ID,
		UserName: from.User.Name,
		Msg:      msg,
	}).Save(from.DB)

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
func (g *GroupStruct) GetHistory10(ctx *Context) (out []string) {
	data, err := database.GroupTable{}.GetLast10(ctx.DB)
	if err != nil {
		log.Println(err)
		return nil
	}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].ID > data[j].ID
	})
	out = make([]string, 0, 10)
	for _, v := range data {
		out = append(out, buf.NewBuf().WriteString("[").
			WriteString(v.UserName).
			WriteString("]: ").
			WriteString(v.Msg).ToString(true))
	}
	return out
}
