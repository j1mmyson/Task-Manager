package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/j1mmyson/reviewList/models"
)

// GET /api/:user_id/card
// Return card list by user_id
func GetCards(c *gin.Context) {
	var cards []models.List
	userID := c.Param("user_id")

	if err := models.DB.Where("user_id = ?", userID).Find(&cards).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cards)
}

// POST /api/card
// data: { user_id, title, state, content, date}
// Create card and return created card
func CreateCard(c *gin.Context) {
	var card models.List
	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Create(&card).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, card)
}

// DELETE /api/card/:id
// Delete card and return boolean on success or failure
func DeleteCard(c *gin.Context) {
	cardID := c.Param("id")

	if err := models.DB.Delete(&models.List{}, cardID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete complete"})
}

// PUT /api/card/:id
// Edit card and return edited card
func EditCard(c *gin.Context) {
	var input models.List
	var card models.List
	cardID := c.Param("id")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("id = ?", cardID).Find(&card).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	copyCard(&input, &card)

	if err := models.DB.Save(&card).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Edit complete"})
}

func copyCard(src, dst *models.List) {
	dst.Title = src.Title
	dst.Content = src.Content
	dst.State = src.State
}
