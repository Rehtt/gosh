package client

import (
	"github.com/Rehtt/Kit/buf"
	"github.com/fatih/color"
)

type Help struct {
}

func (h *Help) Run(ctx *Context, src string) {
	b := buf.NewBuf()
	for k, v := range commands {
		b.WriteColor("/"+k, color.FgHiRed).
			WriteString(" ").
			WriteColor(v.Help(), color.FgHiGreen).WriteByte('\n')
	}
	ctx.Term.Write(b.ToBytes())
}

func (h *Help) Key() string {
	return "help"
}

func (h *Help) Help() string {
	return "show shelp"
}

func init() {
	RegisterCmd(&Help{})
}
