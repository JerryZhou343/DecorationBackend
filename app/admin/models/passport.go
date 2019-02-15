package models

import (
	"github.com/mfslog/DecorationBackend/db"
	"log"
	"time"
)

type Auth struct {
	ID       int
	UserName string
	Salt     string
	Passwd   string    `xorm:"varchar(200)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

func CheckPassport(username, passwd string) bool {
	engine := db.DB()

	has, err := engine.Where("user_name", username).Where("passwd", passwd).Exist()
	log.Printf("%v", err)
	if has {
		return true
	} else {
		return false
	}
}
