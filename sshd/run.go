package sshd

import (
	"fmt"
	"github.com/Rehtt/RehttKit/buf"
	"github.com/Rehtt/gosh/client"
	"github.com/Rehtt/gosh/conf"
	"github.com/Rehtt/gosh/database"
	"github.com/Rehtt/gosh/util"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
	"gorm.io/gorm"
	"log"
	"net"
)

func Run(conf *conf.Conf, db *gorm.DB) error {
	config := &ssh.ServerConfig{
		KeyboardInteractiveCallback: authKeyboard(db),
	}

	for k, v := range privateKeyMap {
		private, err := ssh.ParsePrivateKey(v)
		if err != nil {
			return fmt.Errorf("无效%s证书: %v", k, err)
		}
		config.AddHostKey(private)
	}

	listener, err := net.Listen("tcp", conf.Server.Addr)
	if err != nil {
		return fmt.Errorf("监听地址失败 %v", err)
	}
	fmt.Println("start", conf.Server.Addr)
	//
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listener.Accept", err)
			continue
		}
		go func(conn net.Conn) {
			sConn, cha, req, err := ssh.NewServerConn(conn, config)
			if err != nil {
				log.Println("ssh.NewServerConn", err)
				return
			}
			go ssh.DiscardRequests(req)
			handleChannels(sConn, cha, db)
		}(conn)
	}
}
func handleChannels(sshConn *ssh.ServerConn, channels <-chan ssh.NewChannel, db *gorm.DB) {
	user := database.UserTable{}.FormJson(sshConn.Permissions.Extensions["userJson"])

	// 新建客户端
	ctx := client.NewClient(user, sshConn, db)
	ctx.Group = client.Group
	ctx.Group.Login(ctx)
	go func() {
		sshConn.Wait()
		ctx.Group.Logout(ctx)
	}()

	prompt := buf.NewBuf().WriteString("[").WriteString(user.Name).WriteString("] > ")

	for ch := range channels {
		if t := ch.ChannelType(); t != "session" {
			ch.Reject(ssh.UnknownChannelType, fmt.Sprintf("unknown channel type: %s", t))
			continue
		}
		channel, requests, err := ch.Accept()
		if err != nil {
			log.Printf("Could not accept channel: %v", err)
			continue
		}
		defer channel.Close()

		ctx.Term = term.NewTerminal(channel, prompt.ToString())

		for req := range requests {
			var width, height int
			var ok bool
			switch req.Type {
			case "shell":
				if ctx.Term != nil {
					go ctx.HandleShell(channel)
					ok = true
				}
			case "pty-req":
				//通过如下消息可以让服务器为Session分配一个虚拟终端
				//当客户端的终端窗口大小被改变时，或许需要发送这个消息给服务器。
				width, height, ok = util.ParsePtyRequest(req.Payload)
				if ok {
					err := ctx.Resize(width, height)
					ok = err == nil
				}
			case "window-change":
				width, height, ok = util.ParseWinchRequest(req.Payload)
				if ok {
					err := ctx.Resize(width, height)
					ok = err == nil
				}
			case "exec":
				// ssh rehtt@rehtt.com -p2222 help
			case "env":
				//在shell或command被开始时之后，或许有环境变量需要被传递过去。然而在特权程序里不受控制的设置环境变量是一个很有风险的事情，
				//所以规范推荐实现维护一个允许被设置的环境变量列表或者只有当sshd丢弃权限后设置环境变量。
				log.Print(string(req.Payload))
			case "subsystem":
				//

			default:
				log.Println(req.Type, string(req.Payload))
			}
			if req.WantReply {
				req.Reply(ok, nil)
			}
		}
	}
}
