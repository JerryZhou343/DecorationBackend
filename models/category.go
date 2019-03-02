package models

import (
	"github.com/mfslog/DecorationBackend/db"
)

//TCategory 分类表
type TCategory struct {
	ID         int    `xorm:"pk autoincr"`
	Name       string `xorm:"varchar(64)"`
	ParentID   int    `xorm:"'parent_id'"`
	Priority   int
	State      int    `xorm:"default 1"`
	Remark     string `xorm:"varchar(200)"`
	Created    int    `xorm:"created"`
	Updated    int    `xorm:"updated"`
	OperatorID int    `xorm:"'operator_id'"`
}

//GetChildCategoryByParentID 查询以该分类为父分类的子分类
func GetChildCategoryByParentID(pid int) ([]TCategory, error) {
	engine := db.DB()
	result := []TCategory{}
	err := engine.Where("parent_id=?", pid).Find(&result)
	return result, err
}

//GetCategoryByID 查询分类信息
func GetCategoryByID(id int) (*TCategory, error) {
	engine := db.DB()
	result := TCategory{}
	_, err := engine.Where("id=?", id).Get(&result)

	return &result, err
}

//InsertCategory 插入一条分类信息
func InsertCategory(category *TCategory) error {
	engine := db.DB()
	cnt, err := engine.InsertOne(category)
	if err != nil || cnt > 0 {
		return err
	}

	return nil

}

//UpdateCategoryInfo 更新一条分类信息
func UpdateCategoryInfo(id int, category *TCategory) error {
	engine := db.DB()

	cnt, err := engine.ID(id).Update(category)
	//TODO:区分失败的原因
	if cnt == 0 || err != nil {
		return err
	}

	return nil
}

//DelCategory 删除一条分类信息
func DelCategory(id int) error {
	engine := db.DB()
	tmp := TCategory{
		State: 0,
	}
	_, err := engine.ID(id).Update(&tmp)

	return err
}
