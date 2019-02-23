package models

import (
	"github.com/mfslog/DecorationBackend/db"
)

type TPicCategory struct {
	Id           int `xorm:"pk autoincr"`
	PicId        int
	CategoryId   int
	CategoryName string `xorm:"-"`
	Created      int    `xorm:"created"`
	Updated      int    `xorm:"updated"`
	State        int    `xorm:"tinyint not nul default 1"`
}

func InsertOnePicCategory(picId, categoryId int) error {
	engine := db.DB()
	tmp := TPicCategory{
		PicId:      picId,
		CategoryId: categoryId,
	}
	_, err := engine.InsertOne(&tmp)
	return err
}

func DelPicCategory(picId, categoryId int) error {
	engine := db.DB()
	tmp := TPicCategory{
		PicId:      picId,
		CategoryId: categoryId,
		State:      0,
	}

	_, err := engine.Update(tmp)
	return err
}

//todo:待补充完善
func GetPicCategory(picId int) (*[]TPicCategory, error) {
	//engine := db.DB()

	//err := engine.Query("sele")

	return nil, nil
}
