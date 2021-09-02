package main

import (
	"flag"
	"gosh/constants"
	"gosh/server"
	"gosh/terminal"
)

var (
	udpPort = flag.Int("port", constants.MinPort, "udp server port")
	secret  = flag.String("secret", "", "start server secret")
	ip      = flag.String("ip", "", "remote ssh ip")
	port    = flag.String("p", "22", "remote ssh port")
	user    = flag.String("u", "", "remote ssh user name")
)

func main() {
	flag.Parse()
	if *secret != "" {
		server.Start(*udpPort, []byte(*secret))
	} else {
		terminal.Run(*ip, *port, *user)
	}

}
