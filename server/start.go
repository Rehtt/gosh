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
	w      *sync.WaitGroup
	dog    *utils.Dog
)

func Start(por int, sec []byte) {
	port = por
	secret = sec
	w.Add(1)
	go func() {
		if port <= constants.MaxPort && port >= constants.MinPort {
			for ; port <= constants.MaxPort; port++ {
				udp()
			}
		} else {
			udp()
		}
		fmt.Println("error")
	}()
	w.Wait()
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

	// watchdog
	dog = utils.NewWatchDog(w)

	cmd := make([]byte, 540)
	for {
		index, addr, err := connect.ReadFromUDP(cmd)
		if err != nil {
			fmt.Println(err)
			break
		}
		//
		dog.FeedDog()
		run(cmd[:index], connect, addr)
	}
}
func run(cmd []byte, connect *net.UDPConn, addr *net.UDPAddr) {
	out := ""
	defer func() {
		if out != "" {
			connect.WriteToUDP([]byte(out), addr)
			w.Done()
		}
	}()
	cmd, err := utils.GoshDecrypt(cmd, secret)
	if err != nil {
		fmt.Println(err)
		out = "error secret"
		return
	}
	utils.Cmd(string(cmd), func(out string) (exit bool) {
		o, err := utils.GoshEncrypt([]byte(out), secret)
		if err != nil {
			out = "secret does not meet the regulations"
			return true
		}
		connect.WriteToUDP(o, addr)
		return false
	})
}
