package main

import (
	_ "api-sambasku/conf"
	"api-sambasku/ctrl"
	"api-sambasku/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var appName string
var port string

func initConfig() {
	appName = viper.GetString("appName")
	port = ":" + viper.GetString("server.port")
}

func main() {
	r := gin.Default()

	initConfig()
	db.Connect()
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
