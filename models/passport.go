package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/mfslog/DecorationBackend/db"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

//TAuth 用户信息表
type TAuth struct {
	ID        int64  `xorm:"'id' pk autoincr not null"`
	UserName  string `xorm:"varchar(64) not null unique"`
	Password  string `xorm:"varchar(200)"`
	SlatID    int64  `xorm:"'slat_id' not null"`
	CreatedAt int    `xorm:"created_at"`
	UpdatedAt int    `xorm:"updated_at"`
}

//TSlat 用户盐表
type TSlat struct {
	ID        int64  `xorm:"'id' pk autoincr not null"`
	Slat      string `xorm:"varchar(6) not null"`
	CreatedAt int    `xorm:"created_at"`
	UpdatedAt int    `xorm:"updated_at"`
}

// CheckPassport 校验用户登录信息是否正确
func CheckPassport(username, passwd string) (bool, int64) {
	engine := db.DB()
	auth := TAuth{}
	slat := TSlat{}
	has, _ := engine.Where("user_name=?", username).Get(&auth)
	if has {
		has, _ = engine.Where("id=?", auth.SlatID).Get(&slat)
		if has {
			password := md5.Sum([]byte(slat.Slat + passwd))
			if fmt.Sprintf("%x", password[:]) == auth.Password {
				logrus.Infof("user [%v] check password success ", username)
				return true, auth.ID
			}
			logrus.Errorf("auth password failed for user: [%v]", username)
			return false, 0
		}
		logrus.Errorf("not found slat for user [%v]", username)
		return false, 0
	}
	logrus.Errorf("not found user for user [%v]", username)
	return false, 0
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//RandStringRunes  生成随机字符串
//https://colobu.com/2018/09/02/generate-random-string-in-Go/
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//ChangePassword 修改用户密码
func ChangePassword(username, oldPassword, newPassword string) (err error) {
	flag, uid := CheckPassport(username, oldPassword)
	if !flag {
		return errors.New("not found user " + username)
	}

	slatStr := RandStringRunes(6)

	engine := db.DB()
	slat := &TSlat{
		Slat: slatStr,
	}
	password := md5.Sum([]byte(slat.Slat + newPassword))
	passwordStr := fmt.Sprintf("%x", password[:])
	_, err = engine.InsertOne(slat)

	if err != nil {
		logrus.Errorf("insert slat failed [%v]", err)
		return err
	}

	auth := TAuth{
		Password: passwordStr,
		SlatID:   slat.ID,
	}

	engine.ID(uid).Update(&auth)
	return nil
}
