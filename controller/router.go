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
	// ID      uint        `json:"id" gorm:"primary_key"`
	UserID  string    `form: "user_id" json:"user_id"`
	Title   string    `form: "title" json:"title"`
	State   string    `form: "state" json:"state"`
	Content string    `form: "content" json:"content"`
	Date    time.Time `form: "date" json:"date"`
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

	input.UserID = c.Request.PostFormValue("user_id")
	input.Title = c.Request.PostFormValue("title")
	input.Content = c.Request.PostFormValue("content")
	input.State = c.Request.PostFormValue("state")
	input.Date = time.Now()

	if input.Title == "" || input.Content == "" {
		c.Redirect(http.StatusSeeOther, "/dashboard")
		return
	}

	list := models.List{UserID: input.UserID, Title: input.Title, State: input.State, Date: input.Date, Content: input.Content}
	models.DB.Create(&list)

	c.Redirect(http.StatusSeeOther, "/dashboard")
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

// POST /delete/:id
// Delete card by card id.
func DeleteListById(c *gin.Context) {
	var list models.List

	if err := models.DB.Where("id = ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}

	models.DB.Delete(&list)
	c.Redirect(http.StatusSeeOther, "/dashboard")
}

// POST /edit/:id
// Edit card by card id.
func EditListById(c *gin.Context) {
	var list models.List

	if err := models.DB.First(&list, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}
	list.Title = c.PostForm("title")
	list.Content = c.PostForm("content")
	models.DB.Save(&list)

	c.Redirect(http.StatusSeeOther, "/dashboard")
}

// GET /
// Rendering login page.
func LogInPage(c *gin.Context) {
	if time.Now().Sub(models.DbSessionCleaned) > (time.Second * 30) {
		go models.CleanSessions()
	}

	if IsAleadyLogIn(c) {
		c.Redirect(http.StatusSeeOther, "/dashboard")
		return
	}

	c.HTML(http.StatusOK, "login.html", nil)
}

// POST /
// Check if it is a valid login and redirect user.
func LogIn(c *gin.Context) {
	if time.Now().Sub(models.DbSessionCleaned) > (time.Second * 30) {
		go models.CleanSessions()
	}

	if IsAleadyLogIn(c) {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	var u models.User
	var err error

	if u, err = ReadUser(c); err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{"error": err.Error()})
		return
	}

	sID := uuid.New()
	c.SetCookie("session", sID.String(), CookieDuration, "", "", false, true)
	models.CreateSession(sID.String(), u.ID)

	c.Redirect(http.StatusSeeOther, "/dashboard")
}

// GET /signup
// Rendering signup page.
func SignUpPage(c *gin.Context) {
	if IsAleadyLogIn(c) {
		c.Redirect(http.StatusSeeOther, "/dashboard")
		return
	}
	c.HTML(http.StatusOK, "signup.html", nil)
}

// POST /signup
// Create user account according to user input and check whether it is valid or not.
func SignUp(c *gin.Context) {
	var u models.User
	var err error
	u.ID = c.PostForm("id")
	bs, _ := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.MinCost)
	u.Password = string(bs[:])
	u.Name = c.PostForm("name")

	if err = models.DB.Create(&u).Error; err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"error":    "id already exists.",
			"id":       u.ID,
			"name":     u.Name,
			"password": c.PostForm("password"),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}

// GET /dashboard
// Rendering dashboard page according to session cookie.
func DashBoardPage(c *gin.Context) {
	if !IsAleadyLogIn(c) {
		c.Redirect(http.StatusSeeOther, "/")
	}
	var err error

	sessionValue, err := c.Cookie("session")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	uid, err := models.GetUserIdFromSession(sessionValue)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	data := models.GetCards(uid)

	c.HTML(http.StatusOK, "dashboard.html", data)
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
