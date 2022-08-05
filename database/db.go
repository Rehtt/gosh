package database

import (
	"github.com/Rehtt/gosh/conf"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

func Init(c *conf.Conf) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(c.Database.Path))
	if err != nil {
		log.Panicln(err)
	}
	autoMigrate(db)
	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&UserTable{},
		&GroupTable{},
		&HackTable{},
	)
}
