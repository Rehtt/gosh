package main

import (
	"fmt"
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
	//		out, err := utils.GoshDecrypt(data[:index], []byte("ux0xw7fEbdecurxrWBu2LaSYOVNPdXiU"))
	//		fmt.Println(string(out), err)
	//		x.Done()
	//	}
	//}()
	//
	//conn.Write([]byte{252, 167, 128, 118, 136, 122, 41, 190, 235, 213, 50, 78, 196, 183, 217, 121})
	//x.Wait()
	//time.Sleep(1 * time.Second)

	//a, _ := utils.Cmd(`ssh r@127.0.0.1 "echo SDASDASqwe"`, func(out string) (exit bool) {
	//	fmt.Println(out)
	//	return false
	//})
	//a.Wait()
	//o,_:=utils.AesEncrypt([]byte("ux0qwe"),[]byte("ux0xw7fEbdecurxrWBu2LaSYOVNPdXiU"))
	//s:=base64.StdEncoding.EncodeToString(o)
	//d,err:=base64.StdEncoding.DecodeString("U2FsdGVkX19ZwekFsOhduk4rdeL4r0c31oBDsl4qKWY=")
	//fmt.Println(string(d),err)
	//fmt.Println(o)
	//fmt.Println(utils.AesDecrypt(o,[]byte("ux0xw7fEbdecurxrWBu2LaSYOVNPdXiU")))

	//fmt.Println(string(byte(64)), string(byte(39)), string(byte(46)))
	t := time.NewTimer(5 * time.Second)
	a := time.Now()
	go func() {
		//for{
		select {
		case <-t.C:
			fmt.Println("ww")
			fmt.Println(time.Now().Sub(a))
			break
		}
		//}
	}()
	time.Sleep(1 * time.Second)
	t.Reset(10 * time.Second)
	fmt.Println(time.Now().Sub(a))
	select {}
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
