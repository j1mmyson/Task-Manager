package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/j1mmyson/reviewList/models"
	"golang.org/x/crypto/bcrypt"
)

type CreateListInput struct {
	ID      uint        `json:"id" gorm:"primary_key"`
	UserID  string      `json:"user_id" gorm:"size:191"`
	Title   string      `json:"title"`
	State   string      `json:"state"`
	Content string      `json:"content"`
	User    models.User `gorm:"foreignKey:UserID"`
}

const CookieDuration int = 1800

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

	if time.Now().Sub(models.DbSessionCleaned) > (time.Second * 30) {
		go models.CleanSessions()
	}

	if IsAleadyLogIn(c) {
		// 이미 로그인이 되어있다면 세션을 참조하여 해당 유저의 대시보드로 리다이렉트 해주자
		c.Redirect(http.StatusSeeOther, "/dashboard")
		return
	}
	c.HTML(http.StatusOK, "login.html", nil)
}

// POST /login
// Process login
func LogIn(c *gin.Context) {

	if time.Now().Sub(models.DbSessionCleaned) > (time.Second * 30) {
		go models.CleanSessions()
	}

	if IsAleadyLogIn(c) {
		// 이미 로그인이 되어있다면 세션을 참조하여 해당 유저의 대시보드로 리다이렉트 해주자
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	var u *models.User
	u, err := ReadUser(c)

	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{"error": err.Error()})
		return
	}
	// 쿠키 생성
	sID := uuid.New()
	c.SetCookie("session", sID.String(), CookieDuration, "", "", false, true)

	// 세션 생성
	models.CreateSession(sID.String(), u.ID)

	// 대시보드 렌더링
	c.Redirect(http.StatusSeeOther, "/dashboard")

	// 세션 생성 후 유저 대시보드로 리다이렉트
}

// GET /signup
// SignUp Page
func SignUpPage(c *gin.Context) {

	if IsAleadyLogIn(c) {
		c.Redirect(http.StatusSeeOther, "/dashboard")
		return
	}
	c.HTML(http.StatusOK, "signup.html", nil)
}

// POST /signup
// process signup
func SignUp(c *gin.Context) {
	var u models.User
	var err error
	u.ID = c.PostForm("id")
	bs, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.MinCost)
	u.Password = string(bs[:])
	u.Name = c.PostForm("name")

	if err != nil {

	}

	if err := models.DB.Create(&u).Error; err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"error":    "id already exists.",
			"id":       u.ID,
			"name":     u.Name,
			"password": c.PostForm("password"),
		})
	}
	c.Redirect(http.StatusSeeOther, "/")

}

func DashBoardPage(c *gin.Context) {
	if !IsAleadyLogIn(c) {
		c.Redirect(http.StatusSeeOther, "/")
	}
	// 쿠키에서 세션 밸류값 즉 session_id 값을 받아온다.
	sessionValue, err := c.Cookie("session")
	if err != nil {
		// 세션쿠키가 없다면 로그인페이지로 리다이렉트
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	// 받아온 session_id값으로 데이터베이스의 session 테이블의 user_id값을 가져온다.
	uid, err := models.GetUserIdFromSession(sessionValue)
	if err != nil {
		// 세션이 없다면 로그인 페이지로 리다이렉트
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	u := models.GetUserFromUserId(uid)

	c.HTML(http.StatusOK, "dashboard.html", u)
}

func LogOut(c *gin.Context) {

	if !IsAleadyLogIn(c) {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	sid, _ := c.Cookie("session")
	models.DeleteSession(sid)

	c.SetCookie("session", sid, -1, "", "", false, true)

	if time.Now().Sub(models.DbSessionCleaned) > (time.Minute * 5) {
		go models.CleanSessions()
	}

	c.Redirect(http.StatusSeeOther, "/")

}
