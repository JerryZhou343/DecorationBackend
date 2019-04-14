package models

import (
	"github.com/mfslog/DecorationBackend/db"
	"github.com/sirupsen/logrus"
	"time"
)

//TCategory 分类表
type TCategory struct {
	ID         int    `xorm:"'id' pk autoincr"`
	Name       string `xorm:"varchar(64)"`
	ParentID   int    `xorm:"'parent_id'"`
	Priority   int
	State      int
	Remark     string `xorm:"varchar(200)"`
	CreatedAt  int    `xorm:"created_at"`
	UpdatedAt  int    `xorm:"updated_at"`
	OperatorID int    `xorm:"'operator_id'"`
}

//GetChildCategoryByParentID 查询以该分类为父分类的子分类
func GetChildCategoryByParentID(pid, limit, offset int) ([]TCategory, error) {
	engine := db.DB()
	result := []TCategory{}
	err := engine.Where("parent_id=?", pid).OrderBy("id").Limit(limit, offset).Find(&result)
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
	category.CreatedAt = int(time.Now().Unix())
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

		logrus.Errorf("update category [%v], error [%v]", id, err)
		return err
	} else {
		logrus.Infof("update category [%v], cnt [%v]", id, cnt)
	}

	return nil
}

//DelCategory 删除一条分类信息
func DelCategory(id int) error {
	engine := db.DB()
	tmp := &TCategory{
		State: 0,
	}
	cnt, err := engine.ID(id).Update(tmp)
	if cnt == 0 || err != nil {
		logrus.Errorf("delete category id[%v] error [%v]", id, err)
		return err
	} else {
		logrus.Infof("delete category [%v], cnt [%v]", id, cnt)
	}

	return nil
}
