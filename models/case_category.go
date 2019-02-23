package models

import "github.com/mfslog/DecorationBackend/db"

type TCaseCategory struct {
	Id         int `xorm:"pk autoincr`
	CaseId     int
	CategoryId int
	Created    int `xorm:"created"`
	Updated    int `xorm:"updated"`
	State      int
}

func GetCategoryByCaseId(id int) (*[]*TCaseCategory, error) {
	engine := db.DB()
	ret := []*TCaseCategory{}
	err := engine.Where("case_id=?", id).Find(&ret)

	return &ret, err
}

func InsertOneCaseCategory(category *TCaseCategory) error {
	engine := db.DB()
	cnt, err := engine.InsertOne(category)
	if cnt != 1 || err != nil {
		return err
	}

	return nil
}

func DelCaseCategoryById(caseId, categoryId int) error {
	engine := db.DB()

	_, err := engine.Where("case_id", caseId).
		Where("category_id", categoryId).
		Update(TCaseCategory{
			State: 0,
		})
	return err
}
