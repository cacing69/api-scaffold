package ctrl

import (
	"api-sambasku/db"
	"api-sambasku/mod"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func IndexUser(c *gin.Context) {
	var users []mod.User

	db.T.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {

	var req = struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	var val string

	if err := c.ShouldBind(&req); err != nil {
		if e, ok := err.(validator.ValidationErrors); ok {
			for _, err := range e {
				switch err.Tag() {
				case "required":
					val = fmt.Sprintf("%s is required",
						err.Field())
				case "email":
					val = fmt.Sprintf("%s is not valid email",
						err.Field())
				case "gte":
					val = fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param())
				case "lte":
					val = fmt.Sprintf("%s value must be lower than %s",
						err.Field(), err.Param())
				}
				break
			}
		}
		c.JSON(400, gin.H{"e": val})
		return
	}

	user := mod.User{Name: req.Name, Email: req.Email, Password: req.Password}
	db.T.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func FindUser(c *gin.Context) {
	var user mod.User

	if err := db.T.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record_not_found"})
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var user mod.User

	if err := db.T.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record_not_found"})
	}

	var req = struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db.T.Model(&user).Updates(req)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var user mod.User

	if err := db.T.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record_not_found"})
	}

	db.T.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
