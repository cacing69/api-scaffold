package config

import (
	"log"

	"github.com/spf13/viper"
)

type Db struct {
	Port int
	Name string
}

type conf struct {
	Db Db
}

var T conf

func init() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()

	if err != nil {
		log.Println(err.Error())
	}

	fill()
}

func fill() {
	T.Db.Port = viper.GetInt("database.port")
	T.Db.Name = viper.GetString("database.name")
}
