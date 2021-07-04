package controller

import (
	"fmt"
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
	fmt.Println(c.ContentType())
	var input CreateListInput

	switch c.ContentType() {
	case "multipart/form-data":
		fmt.Println("ContentType : form data")
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	case "application/json":
		fmt.Println("ContentType: json")
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

func DeleteList(c *gin.Context) {

}

func LogIn(c *gin.Context) {

}

func LogOut(c *gin.Context) {

}
