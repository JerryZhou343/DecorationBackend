package models

import "github.com/mfslog/DecorationBackend/db"

type TCase struct {
	Id          int    `xorm:"pk autoincr"`
	Name        string `xorm:"varchar(64)"`
	Price       int
	Type        int
	OwnerName   string `xorm:"varchar(64)"`
	PhoneNumber string `xorm:"varchar(20)"`
	Addr        string `xorm:"varchar(200)"`
	State       int    `xorm:"tinyint default 1"`
	Created     int    `xorm:"created"`
	Updated     int    `xorm:"updated"`
}

func GetCaseById(id int) (*TCase, error) {
	engine := db.DB()
	result := TCase{}
	_, err := engine.Where("id=?", id).Get(&result)

	return &result, err
}

func InsertOneCase(tCase *TCase) error {
	engine := db.DB()
	cnt, err := engine.InsertOne(tCase)
	if cnt != 1 || err != nil {
		return err
	}

	return nil
}

func DelCaseById(id int) error {
	engine := db.DB()
	_, err := engine.Id(id).Update(TCase{
		State: 0,
	})

	return err
}

func UpdateCase(id int, tcase *TCase) error {
	engine := db.DB()
	_, err := engine.Id(id).Update(tcase)
	return err
}
