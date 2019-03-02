package models

import (
	"crypto/md5"
	"fmt"
	"github.com/mfslog/DecorationBackend/db"
)

//TAuth 用户信息表
type TAuth struct {
	ID       int    `xorm:"'id' pk autoincr not null"`
	UserName string `xorm:"varchar(64) not null unique"`
	Password string `xorm:"varchar(200)"`
	Created  int    `xorm:"created"`
	Updated  int    `xorm:"updated"`
}

//TSlat 用户盐表
type TSlat struct {
	ID      int    `xorm:"pk autoincr not null"`
	Slat    string `xorm:"varchar(6) not null"`
	Created int    `xorm:"created"`
	Updated int    `xorm:"updated"`
}

// CheckPassport 校验用户登录信息是否正确
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
			}
			return false

		}
		return false

	}
	return false

}
