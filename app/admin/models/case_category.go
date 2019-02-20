package models

type TCaseCategory struct {
	id         int `xorm:"pk autoincr`
	CaseId     int
	CategoryId int
	Created    int `xorm:"created"`
	Updated    int `xorm:"updated"`
	state      int
}
