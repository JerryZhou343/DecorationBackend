package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/mfslog/DecorationBackend/config"
)

var defaultEngine *xorm.Engine

func Init() error {
	var err error

	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", config.MySQLUser(), config.MySQLPassword(),
		config.MySQLAddr(), config.MySQLDBName())

	defaultEngine, _ = xorm.NewEngine("mysql", dataSource)
	fmt.Println(dataSource)
	err = defaultEngine.Ping()
	defaultEngine.ShowSQL(true)
	if err != nil {
		return err
	}

	defaultEngine.SetMaxIdleConns(config.MySQLPoolLimit())
	return nil
}

//返回数据库连接
func DB() *xorm.Engine {
	return defaultEngine
}
