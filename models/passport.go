package models

import (
	"crypto/md5"
	"fmt"
	"github.com/mfslog/DecorationBackend/db"
)

type TAuth struct {
	Id       int    `xorm:"pk autoincr not null"`
	UserName string `xorm:"varchar(64) not null unique"`
	Password string `xorm:"varchar(200)"`
	Created  int    `xorm:"created"`
	Updated  int    `xorm:"updated"`
}

type TSlat struct {
	Id      int    `xorm:"pk autoincr not null"`
	Slat    string `xorm:"varchar(6) not null"`
	Created int    `xorm:"created"`
	Updated int    `xorm:"updated"`
}

func CheckPassport(username, passwd string) bool {
	engine := db.DB()
	auth := TAuth{}
	slat := TSlat{}
	has, _ := engine.Where("user_name=?", username).Get(&auth)
	if has {
		has, _ = engine.Where("id=?", auth.Id).Get(&slat)
		if has {
			password := md5.Sum([]byte(slat.Slat + passwd))
			if fmt.Sprintf("%x", password[:]) == auth.Password {
				return true
			} else {
				//fmt.Println(slat.Slat)
				//fmt.Printf("%x", string(password[:]))
				return false
			}
		} else {
			return false
		}

	} else {
		return false
	}
}
