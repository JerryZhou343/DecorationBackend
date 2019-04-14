package models

import "github.com/mfslog/DecorationBackend/db"

//TCase 表结构t_case 描述
type TCase struct {
	ID          int    `xorm:"'id'pk autoincr"`
	Name        string `xorm:"varchar(64)"`
	Price       int
	Type        int
	OwnerName   string `xorm:"varchar(64)"`
	PhoneNumber string `xorm:"varchar(20)"`
	Addr        string `xorm:"varchar(200)"`
	State       int    `xorm:"tinyint default 1"`
	CreatedAt   int    `xorm:"created_at"`
	UpdatedAt   int    `xorm:"updated_at"`
}

//GetCaseByID 通过id获得一个Case数据对象
func GetCaseByID(id int) (*TCase, error) {
	engine := db.DB()
	result := TCase{}
	_, err := engine.Where("id=?", id).Get(&result)

	return &result, err
}

//InsertOneCase 插入一个case对象
func InsertOneCase(tCase *TCase) error {
	engine := db.DB()
	cnt, err := engine.InsertOne(tCase)
	if cnt != 1 || err != nil {
		return err
	}

	return nil
}

//DelCaseByID 通过ID删除对应的case对象
func DelCaseByID(id int) error {
	engine := db.DB()
	_, err := engine.Id(id).Update(TCase{
		State: 0,
	})

	return err
}

//UpdateCaseByID 更新对应ID 的case 数据
func UpdateCaseByID(id int, tcase *TCase) error {
	engine := db.DB()
	_, err := engine.Id(id).Update(tcase)
	return err
}

func GetCases(limit, offset int) (*[]TCase, error) {
	engine := db.DB()
	result := []TCase{}
	err := engine.OrderBy("id").Limit(limit, offset).Find(&result)
	return &result, err
}

func GetCaseByCategoryID(categoryID, limit, offset int) (*[]TCase, error) {
	engine := db.DB()
	result := []TCase{}
	err := engine.Where("type=?", categoryID).OrderBy("id").Limit(limit, offset).Find(&result)
	return &result, err
}
