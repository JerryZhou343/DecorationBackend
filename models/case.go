package models

type TCase struct {
	Id          int    `xorm:"pk autoincr"`
	Name        string `xorm:"varchar(64)"`
	Price       int
	Type        int
	OwnerName   string `xorm:"varchar(64)"`
	PhoneNumber string `xorm:"varchar(20)"`
	Addr        string `xorm:"varchar(200)"`
	Created     int    `xorm:"created"`
	Updated     int    `xorm:"updated"`
}
