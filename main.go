package main

import (
	"flag"
)

var (
	confPath = flag.String("conf", "./conf/conf.ini", "config file")
)

func main() {
	flag.Parse()
	panic(Run(*confPath))
}
