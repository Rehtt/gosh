package main

import (
	"fmt"
	"gosh/utils"
	"math/rand"
	"time"
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

	//a, _ := utils.Cmd(`ssh r@127.0.0.1 "echo SDASDASqwe"`, func(out string) (exit bool) {
	//	fmt.Println(out)
	//	return false
	//})
	//a.Wait()

	fmt.Println(string(utils.AesGenerateKey(utils.AES256)))
}

func Shuffle(slice []byte) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}
