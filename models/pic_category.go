package models

import (
	"github.com/mfslog/DecorationBackend/db"
)

//TPicCategory 图片分类表
type TPicCategory struct {
	ID           int    `xorm:"'id'pk autoincr"`
	PicID        int    `xorm:"'pic_id'"`
	CategoryID   int    `xorm:"'category_id'"`
	CategoryName string `xorm:"-"`
	CreatedAt    int    `xorm:"created_at"`
	UpdatedAt    int    `xorm:"updated_at"`
	State        int    `xorm:"tinyint not nul default 1"`
}

//InsertOnePicCategory 插入一个图片分类
func InsertOnePicCategory(picID, categoryID int) error {
	engine := db.DB()
	tmp := TPicCategory{
		PicID:      picID,
		CategoryID: categoryID,
	}
	_, err := engine.InsertOne(&tmp)
	return err
}

//DelPicCategory 删除一个图片分类
func DelPicCategory(picID, categoryID int) error {
	engine := db.DB()
	tmp := TPicCategory{
		PicID:      picID,
		CategoryID: categoryID,
		State:      0,
	}

	_, err := engine.Update(tmp)
	return err
}

//GetPicCategory 获得一个图片分类信息
func GetPicCategory(picID int) (*[]TPicCategory, error) {
	//engine := db.DB()

	//err := engine.Query("sele")

	return nil, nil
}
