package client

type Command interface {
	Key() string
	Help() string
	Run(ctx *Context, src string)
}

var (
	commands = make(map[string]Command)
)

func RegisterCmd(c Command) {
	commands[c.Key()] = c
}

func ParseCmd(src string) (Command, bool) {
	if len(src) > 0 && src[0] == '/' {
		if c, ok := commands[src[1:]]; ok {
			return c, ok
		}
	}
	return nil, false
}
