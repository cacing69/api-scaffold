package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("mysql", "root:cacing.mysql@tcp(127.0.0.1:3306)/db_dummy?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	database.AutoMigrate(&UserModel{})

	DB = database
}
