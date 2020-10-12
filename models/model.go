package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB 数据库连接
var DB *gorm.DB
var err error

func initDB() {
	DB, err = gorm.Open("mysql", "root:20000927@/domesticDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	// 最大空闲连接数
	DB.DB().SetMaxIdleConns(50)
	// DB.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)
	DB.SingularTable(true)
}
func init() {
	initDB()
}
