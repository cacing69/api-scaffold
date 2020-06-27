package db

import (
	"api-sambasku/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var T *gorm.DB
var Type string
var Host string
var Port string
var Password string
var UserName string
var Name string

func Connect() {
	format := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", UserName, Password, Host, Port, Name)
	log.Println(format)
	database, err := gorm.Open(Type, format)

	if err != nil {
		panic(err.Error())
	}

	database.AutoMigrate(&models.UserModel{})

	T = database
}
