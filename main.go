package main

import (
	"github.com/gin-gonic/gin"
	"github.com/j1mmyson/reviewList/controller"
	"github.com/j1mmyson/reviewList/models"
)

func main() {
	r := gin.Default()

	models.ConnectDB()

	r.GET("/lists", controller.AllLists)
	r.POST("/lists", controller.CreateList)
	r.Run()
}
