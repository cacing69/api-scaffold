package main

import (
	"api-sambasku/ctrl"
	"api-sambasku/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var appName string
var port string

func initConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()

	if err != nil {
		log.Println(err.Error())
	}

	db.Type = viper.GetString("database.type")
	db.Host = viper.GetString("database.host")
	db.Port = viper.GetString("database.port")
	db.Password = viper.GetString("database.password")
	db.UserName = viper.GetString("database.username")
	db.Name = viper.GetString("database.name")
	appName = viper.GetString("appName")
	port = ":" + viper.GetString("server.port")
}

func main() {
	r := gin.Default()

	initConfig()
	database.Connect()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": appName})
	})

	r.GET("/users", ctrl.IndexUser)
	r.POST("/user", ctrl.CreateUser)
	r.GET("/user/:id", ctrl.FindUser)
	r.PATCH("/user/:id", ctrl.UpdateUser)
	r.DELETE("/user/:id", ctrl.DeleteUser)

	r.Run(port)
}
