package conf

import (
	"gopkg.in/ini.v1"
	"log"
)

type Conf struct {
	Database struct {
		Path string `ini:"path"`
	} `ini:"database"`
	Server struct {
		Addr string `ini:"addr"`
	} `ini:"server"`
}

var Salf = new(Conf)

func Parse(path string) *Conf {
	err := ini.MapTo(Salf, path)
	if err != nil {
		log.Panicln(err)
	}
	return Salf
}
