package terminal

import (
	"bufio"
	"fmt"
	"gosh/utils"
	"net"
	"os"
	"strings"
)

var (
	udpPort string
	secret  = utils.AesGenerateKey(utils.AES256)
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
	startSession(ip, udpPort)
}

func startSession(ip, port string) {
	dstAddr, err := net.ResolveUDPAddr("udp", ip+":"+port)
	conn, err := net.DialUDP("udp", nil, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	// send "start" start session
	conn.Write([]byte("start"))
	sessionNumber := byte(0)
	data := make([]byte, 500)
	go func() {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			in := append([]byte{sessionNumber + 1}, input.Bytes()...)
			conn.Write(in)
		}
	}()
	for {
		index, _, _ := conn.ReadFromUDP(data)
		out, err := utils.GoshDecrypt(data[:index], secret)
		if err != nil {
			fmt.Println(err)
			return
		}
		if sessionNumber == 0 && string(out) == "start" {

		} else if sessionNumber < out[0] {
			fmt.Println(string(out[1:]))
		}
	}
}
