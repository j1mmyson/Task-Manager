package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/j1mmyson/reviewList/models"
	"golang.org/x/crypto/bcrypt"
)

func IsAleadyLogIn(c *gin.Context) bool {
	sid, err := c.Cookie("session")
	if err != nil {
		return false
	}
	uid, err := models.GetUserIdFromSession(sid)
	if err != nil {
		// 쿠키는 있는데 세션이 없으면 쿠키 삭제
		return false
	}
	c.SetCookie("session", sid, 1800, "", "", false, true)
	models.UpdateCurrentTime(*models.GetSessionFromUserId(uid))
	return true
}

func ReadUser(c *gin.Context) (*models.User, error) {
	id, pw := c.PostForm("id"), c.PostForm("password")

	var u models.User

	if err := models.DB.First(&u, "id = ?", id).Error; err != nil {
		return nil, errors.New("id doesn`t exists")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw)); err != nil {
		return nil, errors.New("wrong password T^T")
	}

	return &u, nil

}
