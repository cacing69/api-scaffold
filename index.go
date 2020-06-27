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

	r.GET("/tester", action.IndexTester)
	r.POST("/tester", action.CreateTester)
	r.GET("/tester/:id", action.FindTester)
	r.PATCH("/tester/:id", action.UpdateTester)
	r.DELETE("/tester/:id", action.DeleteTester)

	r.Run(port)
}
