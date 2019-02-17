package models

import (
	"github.com/mfslog/DecorationBackend/db"
)

type TCategory struct {
	Id         int    `xorm:"pk autoincr"`
	Name       string `xorm:"varchar(64)"`
	ParentId   int
	Priority   int
	State      int    `xorm:"default 1"`
	Remark     string `xorm:"varchar(200)"`
	Created    int    `xorm:"created"`
	Updated    int    `xorm:"updated"`
	OperatorId int
}

func GetChildCategoryByParentId(pid int) ([]*TCategory, error) {
	engine := db.DB()
	result := []*TCategory{}
	_, err := engine.Where("parent_id=?", pid).Get(&result)

	return result, err
}

func GetCategoryById(id int) (*TCategory, error) {
	engine := db.DB()
	result := TCategory{}
	_, err := engine.Where("id=?", id).Get(&result)

	return &result, err
}

func InsertCategory(category *TCategory) error {
	engine := db.DB()
	cnt, err := engine.InsertOne(category)
	if err != nil || cnt > 0 {
		return err
	}

	return nil

}

func UpdateCategoryInfo(id int, category *TCategory) error {
	engine := db.DB()

	cnt, err := engine.ID(id).Update(category)
	//TODO:区分失败的原因
	if cnt == 0 || err != nil {
		return err
	}

	return nil
}

func DelCategory(id int) error {
	engine := db.DB()
	tmp := TCategory{
		State: 0,
	}
	cnt, err := engine.ID(id).Update(&tmp)
	//TODO:区分失败的原因
	if cnt == 0 || err != nil {
		return err
	}

	return nil
}
