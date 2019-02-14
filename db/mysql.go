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
	dataSource := fmt.Sprintf("%s:%s@%s/%s", config.MySQLUser(), config.MySQLPassword(),
		config.MySQLAddr(), config.MySQLDBName())

	defaultEngine, err = xorm.NewEngine("mysql", dataSource)

	defaultEngine.SetMaxIdleConns(config.MySQLPoolLimit())
	return nil
}