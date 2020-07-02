package main

import (
	_ "api-scaffold/config"
	"api-scaffold/controller"
	_ "api-scaffold/db"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var appName string = "main"
var port string = ":8000"

func main() {
	r := gin.Default()
	// db.Connect()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": appName})
	})

	r.GET("/ping", func(c *gin.Context) {
		t := time.NewTicker(1 * time.Second)
		defer t.Stop()

		for i := 1; i < 15; i++ {
			select {
			case <-t.C:
				log.Println(i)
			case <-c.Request.Context().Done():
				if err := c.Request.Context().Err(); err != nil {
					if strings.Contains(strings.ToLower(err.Error()), "canceled") {
						log.Println("request canceled")
					} else {
						log.Println("unknown error occured.", err.Error())
					}
				}
				return
			}
		}

		c.JSON(200, gin.H{
			"data": "ping",
		})
	})

	r.GET("/tester", controller.IndexTester)
	// r.POST("/tester", controller.StoreTester)
	// r.GET("/tester/:id", controller.FindTester)
	// r.PATCH("/tester/:id", controller.UpdateTester)
	// r.DELETE("/tester/:id", controller.DeleteTester)

	r.Run(port)
}
