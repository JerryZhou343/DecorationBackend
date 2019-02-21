package models

type TPicture struct {
	Id      int    `xorm:"pk autoincr"`
	Name    string `xorm:"varchar(64)"`
	Addr    string `xorm:"varchar(1024)"`
	Created int    `xorm:"created"`
}
