package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/j1mmyson/reviewList/models"
)

type CreateListInput struct {
	ID      uint        `json:"id" gorm:"primary_key"`
	UserID  string      `json:"user_id" gorm:"size:191"`
	Title   string      `json:"title"`
	State   string      `json:"state"`
	Content string      `json:"content"`
	User    models.User `gorm:"foreignKey:UserID"`
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
	list := models.List{UserID: input.UserID, Title: input.Title, State: input.State, Content: input.Content}
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

// GET /
// Login Page
func LogInPage(c *gin.Context) {
	if isAleadyLogIn() {
		// 이미 로그인이 되어있다면 세션을 참조하여 해당 유저의 대시보드로 리다이렉트 해주자
		c.Redirect(http.StatusSeeOther, "/")
	}
	c.HTML(http.StatusOK, "login.html", nil)
}

// GET /signup
// SignUp Page
func SignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

// POST /login
// Process login
func LogIn(c *gin.Context) {
	if isAleadyLogIn() {
		// 이미 로그인이 되어있다면 세션을 참조하여 해당 유저의 대시보드로 리다이렉트 해주자
		c.Redirect(http.StatusSeeOther, "/")
	}

	// 세션 생성 후 유저 대시보드로 리다이렉트
}

func LogOut(c *gin.Context) {
	// 세션 삭제 후 로그인 페이지로 리다이렉트

}
