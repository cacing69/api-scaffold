package action

import (
	"api-sambasku/db"
	"api-sambasku/lib"
	"api-sambasku/mod"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexTester(c *gin.Context) {
	var data []mod.Tester

	db.T.Find(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func CreateTester(c *gin.Context) {

	var req = struct {
		Key   string `json:"name" binding:"required"`
		Value string `json:"email" binding:"required,email"`
	}{}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": lib.V.Req.Validate(err)})
		return
	}

	data := mod.Tester{Key: req.Key, Value: req.Value}

	if err := db.T.Create(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func FindTester(c *gin.Context) {
	var data mod.Tester

	if err := db.T.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"record_not_found"}})
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func UpdateTester(c *gin.Context) {
	var test mod.Tester

	if err := db.T.Where("tester_id = ?", c.Param("id")).First(&test).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"record_not_found"}})
	}

	var req = struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }

	db.T.Model(&test).Updates(req)

	c.JSON(http.StatusOK, gin.H{"data": test})
}

func DeleteTester(c *gin.Context) {
	var data mod.Tester

	if err := db.T.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"record_not_found"}})
	}

	db.T.Delete(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}
