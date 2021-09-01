package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	ip := net.ParseIP("127.0.0.1")
	dstAddr := &net.UDPAddr{IP: ip, Port: 58000}
	conn, err := net.DialUDP("udp", nil, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	x := sync.WaitGroup{}
	x.Add(1)
	go func() {
		data := make([]byte, 500)
		for {
			index, _, _ := conn.ReadFromUDP(data)
			fmt.Println(string(data[:index]))
			x.Done()
		}
	}()

	conn.Write([]byte("123"))
	fmt.Println("qwe")
	x.Wait()
	time.Sleep(1 * time.Second)

	//c := utils.Cmd("echo 123&&sleep 5s&&echo 234")
	//c.OutPut(func(out string)bool {
	//	fmt.Println(out)
	//	return true
	//})
	//fmt.Println(c.Run())
}
