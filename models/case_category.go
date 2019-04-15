package models

import "github.com/mfslog/DecorationBackend/db"

//TCaseCategory 案例分类表定义
type TCaseCategory struct {
	ID         int `xorm:"'id' pk autoincr"`
	CaseID     int `xorm:"'case_id'"`
	CategoryID int `xorm:"'category_id'"`
	CreatedAt  int `xorm:"created_at"`
	UpdatedAt  int `xorm:"updated_at"`
	State      int
}

//GetCategoryByCaseID 通过case ID 查询对应case的分类信息
func GetCategoryByCaseID(id int) (*[]*TCaseCategory, error) {
	engine := db.DB()
	ret := []*TCaseCategory{}
	err := engine.Where("case_id=?", id).Find(&ret)

	return &ret, err
}

//InsertOneCaseCategory 插入一个case分类
func InsertOneCaseCategory(category *TCaseCategory) error {
	engine := db.DB()
	cnt, err := engine.InsertOne(category)
	if cnt != 1 || err != nil {
		return err
	}

	return nil
}

// DelCaseCategoryByID 删除对应case和对应的category
func DelCaseCategoryByID(caseID, categoryID int) error {
	engine := db.DB()

	_, err := engine.Where("case_id", caseID).
		Where("category_id", categoryID).
		Update(TCaseCategory{
			State: 0,
		})
	return err
}
