package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 连接mysql，user:password@(localhost)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	db, err := gorm.Open("mysql", "root:Dwh001107!@tcp(127.0.0.1:3306)/gormDB1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{1, "dxp", "man", "跑步"}
	u2 := UserInfo{2, "lxz", "woman", "游泳"}
	// create
	db.Create(&u1)
	db.Create(&u2)

	// select
	u := db.First(&u1)
	fmt.Printf("u=%v\n", u)

}
