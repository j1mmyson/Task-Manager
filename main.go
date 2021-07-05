package main

import (
	"github.com/gin-gonic/gin"
	"github.com/j1mmyson/reviewList/controller"
	"github.com/j1mmyson/reviewList/models"
)

func main() {
	r := gin.Default()

	models.ConnectDB()

	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", controller.IndexPage)
	r.GET("/lists", controller.AllLists)
	r.POST("/lists", controller.CreateList)
	r.GET("/lists/:user", controller.FindListByUserName)
	r.POST("/lists/delete/:id", controller.DeleteListById)
	r.Run()
}
