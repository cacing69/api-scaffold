package controller

import (
	"api-scaffold/db"
	"api-scaffold/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexTester(c *gin.Context) {
	var data []model.Tester
	var count int64

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "20")

	_page, _ := strconv.Atoi(page)
	_limit, _ := strconv.Atoi(limit)

	_offset := (_page - 1) * _limit

	_id := c.Query("id")
	_key := c.Query("key")

	q := db.T.WithContext(c).Model(&model.Tester{})

	if len(_id) > 0 {
		q = q.Where("tester_id = ?", _id)
	}

	if len(_key) > 0 {
		q = q.Where("tester_key like ?", "%"+_key+"%")
	}

	q.Limit(_limit).Offset(_offset).Find(&data)
	q.Count(&count)

	meta := gin.H{
		"total":     count,
		"page":      _page,
		"per_page":  _limit,
		"last_page": 1,
	}

	c.JSON(http.StatusOK, gin.H{"data": data, "meta": meta})
}

// func StoreTester(c *gin.Context) {

// 	var req = struct {
// 		Key   string `json:"name" binding:"required"`
// 		Value string `json:"email" binding:"required,email"`
// 	}{}

// 	if err := c.ShouldBind(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"errors": lib.V.Req.Validate(err)})
// 		return
// 	}

// 	data := model.Tester{Key: req.Key, Value: req.Value}

// 	if err := db.T.Create(&data).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"errors": []string{err.Error()}})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": data})
// }

// func FindTester(c *gin.Context) {
// 	var data model.Tester

// 	if err := db.T.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"record_not_found"}})
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": data})
// }

// func UpdateTester(c *gin.Context) {
// 	var test model.Tester

// 	if err := db.T.Where("tester_id = ?", c.Param("id")).First(&test).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"record_not_found"}})
// 	}

// 	var req = struct {
// 		Name     string `json:"name"`
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}{}

// 	// if err := c.ShouldBindJSON(&req); err != nil {
// 	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	// }

// 	db.T.Model(&test).Updates(req)

// 	c.JSON(http.StatusOK, gin.H{"data": test})
// }

// func DeleteTester(c *gin.Context) {
// 	var data model.Tester

// 	if err := db.T.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"record_not_found"}})
// 	}

// 	db.T.Delete(&data)

// 	c.JSON(http.StatusOK, gin.H{"data": data})
// }
