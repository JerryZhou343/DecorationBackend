package models

import "github.com/mfslog/DecorationBackend/db"

type TPicture struct {
	Id      int `xorm:"pk autoincr"`
	CaseId  int
	Name    string `xorm:"varchar(64)"`
	Remark  string `xorm:"varchar(1024)"`
	Addr    string `xorm:"varchar(1024)"`
	State   int    `xorm:not null default 1`
	Created int    `xorm:"created"`
	Updated int    `xorm:"updated"`
}

func InsertOnePicture(pic *TPicture) error {
	engine := db.DB()
	_, err := engine.InsertOne(pic)
	return err
}

func DelOnePictureById(id int) error {
	engine := db.DB()
	tmp := TPicture{
		State: 0,
	}
	_, err := engine.Where("id = ?", id).Update(tmp)
	return err
}

func UpdateOnePicture(id int, pic *TPicture) error {
	engine := db.DB()
	_, err := engine.Where("id = ?", id).Update(pic)
	return err
}

func GetPictureById(id int) (*TPicture, error) {
	engine := db.DB()
	tmp := &TPicture{}
	_, err := engine.Where("id = ?", id).Get(tmp)

	return tmp, err
}
