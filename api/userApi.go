package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/j1mmyson/reviewList/models"
	"golang.org/x/crypto/bcrypt"
)

// GET /api/user
// Return all users
func ShowUserList(c *gin.Context) {
	var users []models.User

	if err := models.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /api/user/:id
// Return user information ex) id, pw, name, created time ..
func GetUser(c *gin.Context) {
	var u models.User
	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).Find(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

// POST /api/user
// Return boolean value on success or failure
func CreateUser(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bs, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	u.Password = string(bs[:])

	if err := models.DB.Create(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

// DELETE /api/user
func DeleteUser(c *gin.Context) {
	var input models.User
	var u models.User
	var list models.List

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.First(&u, "id = ?", input.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id doesn`t exists"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password T^T"})
		return
	}

	if err := models.DB.Where("user_id = ?", u.ID).Delete(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Delete(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete user complete !"})
}
