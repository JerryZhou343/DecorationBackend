package models

type TCaseCategory struct {
	Id         int `xorm:"pk autoincr`
	CaseId     int
	CategoryId int
	Created    int `xorm:"created"`
	Updated    int `xorm:"updated"`
	State      int
}
