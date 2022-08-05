package client

import (
	"encoding/base64"
	"fmt"
	"github.com/Rehtt/RehttKit/buf"
	"github.com/fatih/color"
	"strings"
)

type Base64 struct{}

const (
	base64Table = "IJjSKLMNO567PQX12RVW3YZaDEF8bcdefghiABCHlkTUmnopqrxyz04stuvwG9+/"
)

func (b *Base64) Key() string {
	return "base64"
}

func (b *Base64) Help() string {
	return "[-d/-e]解码/编码 [a]混淆, eg: /base64 -ae 123 ; /base64 -ad PWOy"
}

func (b *Base64) Run(ctx *Context, src string) {
	split := strings.SplitN(src, " ", 3)
	fmt.Println(split)
	if len(split) >= 3 && split[1][0] == '-' {
		arge := split[1]

		encoding := base64.StdEncoding
		for i := 1; i < len(arge); i++ {
			switch arge[i] {
			case 'd':
				out, err := encoding.DecodeString(split[2])
				if err != nil {
					ctx.Term.Write(buf.NewBuf().WriteColor(err, color.FgHiRed).WriteByte('\n').ToBytes(true))
				}
				ctx.Term.Write(buf.NewBuf().WriteColor(string(out), color.FgHiGreen).WriteByte('\n').ToBytes(true))
			case 'e':
				out := encoding.EncodeToString([]byte(split[2]))
				ctx.Term.Write(buf.NewBuf().WriteColor(out, color.FgHiGreen).WriteByte('\n').ToBytes(true))
			case 'a':
				encoding = base64.NewEncoding(base64Table)
			}
		}
		return
	}
	ctx.Term.Write([]byte("error"))
}

func init() {
	RegisterCmd(&Base64{})
}
