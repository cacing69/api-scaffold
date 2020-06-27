package main

import (
	"api-sambasku/action"
	_ "api-sambasku/conf"
	"api-sambasku/db"

	"github.com/gin-gonic/gin"
)

var appName string = "main"
var port string = ":8000"

func main() {
	r := gin.Default()
	db.Connect()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": appName})
	})

	r.GET("/users", action.IndexUser)
	r.POST("/user", action.CreateUser)
	r.GET("/user/:id", action.FindUser)
	r.PATCH("/user/:id", action.UpdateUser)
	r.DELETE("/user/:id", action.DeleteUser)

	r.Run(port)
}
