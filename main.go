package main

import (
	"flag"
	"gosh/constants"
	"gosh/server"
	"gosh/terminal"
	"sync"
)

var (
	udpPort = flag.Int("port", constants.MinPort, "udp server port")
	secret  = flag.String("secret", "", "start server secret")
	w       sync.WaitGroup
)

func main() {
	flag.Parse()
	if *secret != "" {
		w.Add(1)
		go server.Start(*udpPort, []byte(*secret), &w)
		w.Wait()
	} else {
		// todo start terminal
		terminal.Run()
	}

}
