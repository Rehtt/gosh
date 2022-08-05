package database

import "gorm.io/gorm"

type HackTable struct {
	gorm.Model
	UserName      string
	Password      string
	Addr          string
	ClientVersion string
}

func (HackTable) TableName() string {
	return "hack"
}
func (h *HackTable) Save(db *gorm.DB) {
	db.Save(h)
}
