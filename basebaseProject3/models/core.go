package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 配置模块的时候自动执行init
var DB *gorm.DB
var err error

func init() {
	dsn := "root:yizhihu@tcp(localhost:3305)/recordings?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
