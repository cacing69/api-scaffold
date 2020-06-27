package db

import (
	"api-sambasku/conf"
	"api-sambasku/mod"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var T *gorm.DB
var Type string
var Host string
var Port string
var Password string
var UserName string
var Name string

func init() {
	Type = viper.GetString("database.type")
	Host = viper.GetString("database.host")
	Password = viper.GetString("database.password")
	UserName = viper.GetString("database.username")
}

func Connect() {
	format := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", UserName, Password, Host, conf.T.Db.Port, conf.T.Db.Name)
	log.Println(format)
	database, err := gorm.Open(Type, format)

	if err != nil {
		panic(err.Error())
	}

	database.AutoMigrate(&mod.User{})

	T = database
}
