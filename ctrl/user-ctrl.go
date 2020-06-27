package ctrl

import (
	"api-sambasku/db"
	"api-sambasku/mod"
	"api-sambasku/val"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexUser(c *gin.Context) {
	var users []mod.UserMod

	db.T.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var input val.UserCreateVal

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := mod.UserMod{Name: input.Name, Email: input.Email, Password: input.Password}
	db.T.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func FindUser(c *gin.Context) {
	var user mod.UserMod

	if err := db.T.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record_not_found"})
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var user mod.UserMod

	if err := db.T.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record_not_found"})
	}

	var input val.UserUpdateVal
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db.T.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var user mod.UserMod

	if err := db.T.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record_not_found"})
	}

	db.T.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
