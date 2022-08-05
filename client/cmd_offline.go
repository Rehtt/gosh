package client

type Offline struct{}

func (o *Offline) Key() string {
	return "offline"
}

func (o *Offline) Help() string {
	return "退出"
}

func (o *Offline) Run(ctx *Context, src string) {
	ctx.Conn.Close()
}

func init() {
	RegisterCmd(&Offline{})
}
