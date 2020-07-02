package db

import (
	"api-scaffold/config"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var T *gorm.DB
var dbType string
var host string
var port int
var password string
var userName string
var name string
var debug bool

func init() {
	dbType = viper.GetString("database.type")
	host = viper.GetString("database.host")
	password = viper.GetString("database.password")
	userName = viper.GetString("database.username")
	port = config.T.Db.Port
	name = config.T.Db.Name
	debug = true

	format := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", userName, password, host, port, name)
	db, err := gorm.Open(mysql.Open(format), &gorm.Config{NowFunc: func() time.Time { return time.Now().Local() }})

	if err != nil {
		panic(err.Error())
	}

	if debug {
		db = db.Debug()
	}

	T = db
}
