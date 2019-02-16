package models

import (
	"crypto/md5"
	"fmt"
	"github.com/mfslog/DecorationBackend/db"
	"time"
)

type TAuth struct {
	ID       int       `xorm:"not null primary 'id'"`
	UserName string    `xorm:"varchar(64) not null unique"`
	Password string    `xorm:"varchar(200)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

type TSlat struct {
	ID      int       `xorm:"not null primary 'id'"`
	Slat    string    `xorm:"varchar(6) not null"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func CheckPassport(username, passwd string) bool {
	engine := db.DB()
	auth := TAuth{}
	slat := TSlat{}
	has, _ := engine.Where("user_name=?", username).Get(&auth)
	if has {
		has, _ = engine.Where("id=?", auth.ID).Get(&slat)
		if has {
			password := md5.Sum([]byte(slat.Slat + passwd))
			if fmt.Sprintf("%x", password[:]) == auth.Password {
				return true
			} else {
				fmt.Println(slat.Slat)
				fmt.Printf("%x", string(password[:]))
				return false
			}
		} else {
			return false
		}

	} else {
		return false
	}
}
