package main

import (
	"fmt"
	"gosh/utils"
)

func main() {
	//ip := net.ParseIP("127.0.0.1")
	//dstAddr := &net.UDPAddr{IP: ip, Port: 58000}
	//conn, err := net.DialUDP("udp", nil, dstAddr)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//x := sync.WaitGroup{}
	//x.Add(1)
	//go func() {
	//	data := make([]byte, 500)
	//	for {
	//		index, _, _ := conn.ReadFromUDP(data)
	//		fmt.Println(string(data[:index]))
	//		x.Done()
	//	}
	//}()
	//
	//conn.Write([]byte("123"))
	//fmt.Println("qwe")
	//x.Wait()
	//time.Sleep(1 * time.Second)

	a, _ := utils.Cmd(`echo \\`, func(out string) (exit bool) {
		fmt.Println(out)
		return true
	})
	a.Input.WriteString("123\n")
	a.C.Wait()

}
