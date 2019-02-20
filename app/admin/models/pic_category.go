package models

type TPicCategory struct {
	Id         int `xorm:"pk autoincr"`
	PicId      int
	CategoryId int
	Created    int `xorm:"created"`
	Updated    int `xorm:"updated"`
	State      int
}
