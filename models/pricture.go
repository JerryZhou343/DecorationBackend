package models

import "github.com/mfslog/DecorationBackend/db"

//TPicture 图片表
type TPicture struct {
	ID        int    `xorm:"'id' pk autoincr"`
	CaseID    int    `xorm:"'case_id'"`
	Name      string `xorm:"varchar(64)"`
	Remark    string `xorm:"varchar(1024)"`
	Addr      string `xorm:"varchar(1024)"`
	State     int    `xorm:not null default 1`
	CreatedAt int    `xorm:"created_at"`
	UpdatedAt int    `xorm:"updated_at"`
}

//InsertOnePicture 插入一张图片
func InsertOnePicture(pic *TPicture) error {
	engine := db.DB()
	_, err := engine.InsertOne(pic)
	return err
}

//DelOnePictureByID 通过图片ID 删除一张图片
func DelOnePictureByID(id int) error {
	engine := db.DB()
	tmp := TPicture{
		State: 0,
	}
	_, err := engine.Where("id = ?", id).Update(tmp)
	return err
}

//UpdateOnePicture 通过图片ID 更新图片信息
func UpdateOnePicture(id int, pic *TPicture) error {
	engine := db.DB()
	_, err := engine.Where("id = ?", id).Update(pic)
	return err
}

//GetPictureByID 通过图片ID 获得一张图片信息
func GetPictureByID(id int) (*TPicture, error) {
	engine := db.DB()
	tmp := &TPicture{}
	_, err := engine.Where("id = ?", id).Get(tmp)

	return tmp, err
}
