package database

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

type UserTable struct {
	gorm.Model
	Name          string    `json:"name"`
	LastLoginTime time.Time `json:"last_login_time"`
	QNo           string    `json:"q_no"`
	PublicKey     *string   `json:"public_key"`
	Password      *string   `json:"password"`
}

func (UserTable) TableName() string {
	return "user"
}
func (u *UserTable) ToJson() string {
	out, _ := json.Marshal(u)
	return string(out)
}
func (UserTable) FormJson(j string) (user *UserTable) {
	user = new(UserTable)
	json.NewDecoder(strings.NewReader(j)).Decode(user)
	return
}

func (u *UserTable) Get(db *gorm.DB, username string) *UserTable {
	err := db.Where("name = ?", username).Find(u).Error
	if err != nil {
		log.Println(err)
	}
	return u
}
func (u *UserTable) Save(db *gorm.DB) error {
	return db.Save(u).Error
}
