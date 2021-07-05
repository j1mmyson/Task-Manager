package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/j1mmyson/reviewList/models"
)

type CreateListInput struct {
	User    string `form:"user" json:"user"`
	Title   string `form:"title" json:"title"`
	State   string `form:"state" json:"state"`
	Content string `form:"content" json:"content"`
}

// GET /lists
// Get all lists
func AllLists(c *gin.Context) {
	var lists []models.List
	models.DB.Find(&lists)

	c.JSON(http.StatusOK, gin.H{"data": lists})

}

// POST /lists
// Create List
func CreateList(c *gin.Context) {
	var input CreateListInput

	switch c.ContentType() {
	case "multipart/form-data":
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	case "application/json":
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
			return
		}
	}

	// Create list
	list := models.List{User: input.User, Title: input.Title, State: input.State, Content: input.Content}
	models.DB.Create(&list)

	c.JSON(http.StatusOK, gin.H{"data": list})
}

// GET /lists/:user
// Get lists by user_name
func FindListByUserName(c *gin.Context) {
	var lists []models.List

	if err := models.DB.Where("user = ?", c.Param("user")).Find(&lists).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": lists})
}

// POST /lists/delete/:id
func DeleteListById(c *gin.Context) {
	var list models.List

	if err := models.DB.Where("id = ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}

	models.DB.Delete(&list)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Index Pagez",
	})
}

func LogIn(c *gin.Context) {

}

func LogOut(c *gin.Context) {

}
