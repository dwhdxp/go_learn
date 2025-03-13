package conn

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn/go_web/gorm/config"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		fmt.Printf("connect db fail, err:%v\n", err)
	}
	if Db.Error != nil {
		fmt.Printf("database err, err:%v\n", Db.Error)
	}
}
