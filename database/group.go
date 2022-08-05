package database

import (
	"gorm.io/gorm"
)

type GroupTable struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	Msg      string `json:"msg"`
}

func (GroupTable) TableName() string {
	return "group"
}

func (g *GroupTable) Save(db *gorm.DB) error {
	return db.Save(g).Error
}
func (g GroupTable) GetLast10(db *gorm.DB) (out []GroupTable, err error) {
	err = db.Table(g.TableName()).Limit(10).Order("id desc").Find(&out).Error
	return
}
