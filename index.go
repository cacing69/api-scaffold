package main

import (
	_ "api-scaffold/conf"
	"api-scaffold/controller"
	"api-scaffold/db"

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

	r.GET("/tester", controller.IndexTester)
	r.POST("/tester", controller.CreateTester)
	r.GET("/tester/:id", controller.FindTester)
	r.PATCH("/tester/:id", controller.UpdateTester)
	r.DELETE("/tester/:id", controller.DeleteTester)

	r.Run(port)
}
