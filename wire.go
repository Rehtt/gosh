//go:build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/Rehtt/gosh/conf"
	"github.com/Rehtt/gosh/database"
	"github.com/Rehtt/gosh/sshd"
	"github.com/google/wire"
)

func Run(confPath string) error {
	wire.Build(conf.Parse, database.Init, sshd.Run)
	return nil
}
