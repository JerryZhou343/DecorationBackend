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

func GetChildCategoryByParentId(pid int) []*TCategory {
	engine := db.DB()
	result := []*TCategory{}
	engine.Where("parent_id=?", pid).Get(&result)

	return result
}

func InsertCategory(category *TCategory) bool {
	engine := db.DB()
	ret, err := engine.InsertOne(category)
	if err != nil && ret > 0 {
		return false
	}

	return true

}
