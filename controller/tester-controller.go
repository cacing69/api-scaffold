package controller

import (
	"api-sambasku/db"
	"api-sambasku/lib"
	"api-sambasku/mod"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexTester(c *gin.Context) {
	var data []mod.Tester
	var count int

	// done := make(chan bool)

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "20")

	_page, _ := strconv.Atoi(page)
	_limit, _ := strconv.Atoi(limit)

	_offset := (_page - 1) * _limit

	_id := c.Query("id")
	_key := c.Query("key")

	q := db.T.Model(&mod.Tester{})

	log.Println(_id)
	log.Println(len(_id))

	if len(_id) > 0 {
		q = q.Where("tester_id = ?", _id)
	}

	if len(_key) > 0 {
		q = q.Where("tester_key like ?", "%"+_key+"%")
	}

	// go func() {
	q.Limit(_limit).Offset(_offset).Find(&data)
	q.Count(&count)

	// 	done <- true
	// }()

	// done <- true

	meta := gin.H{
		"total":     count,
		"page":      _page,
		"per_page":  _limit,
		"last_page": 1,
	}
	// select {
	// case <-c.Request.Context().Done():
	// 	// if err := c.Done(); err != nil {
	// 	log.Printf("%#v", c.Err())
	// 	// if strings.Contains(strings.ToLower(err.Error()), "canceled") {
	// 	// 	log.Println("request canceled")
	// 	// } else {
	// 	// 	log.Println("unknown error occured.", err.Error())
	// 	// }
	// 	// }
	// case <-done:
	c.JSON(http.StatusOK, gin.H{"data": data, "meta": meta})
	// }

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
