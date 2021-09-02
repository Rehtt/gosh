package terminal

import (
	"fmt"
	"gosh/utils"
	"net"
	"strings"
)

var (
	udpPort string
	secret  = string(utils.AesGenerateKey(utils.AES256))
)

func Run(ip, p, user string) {
	utils.SSH(fmt.Sprintf(`%s@%s -p %s "/mnt/e/ide/gosh/gosh -secret %s &"`, user, ip, p, secret), func(out string) (exit bool) {
		if strings.Contains(out, "gosh: command not found") {
			fmt.Println("not find gosh")
			return
		}
		if strings.Contains(out, "success") {
			udpPort = strings.Split(out, " ")[1]
			return true
		}
		if strings.Contains(out, "error") {
			fmt.Println("failed to open port")
			return true
		}
		fmt.Println(out)
		return false
	})
	if udpPort == "" {
		return
	}
	openUdp(ip, udpPort)
}

func openUdp(ip, port string) {
	dstAddr, err := net.ResolveUDPAddr("udp", ip+":"+port)
	conn, err := net.DialUDP("udp", nil, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	data := make([]byte, 500)
	for {
		index, _, _ := conn.ReadFromUDP(data)
		fmt.Println(string(data[:index]))
	}
}
