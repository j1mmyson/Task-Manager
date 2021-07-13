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
	r.Static("/web/static", "./web/static")

	r.GET("/", controller.LogInPage)
	r.POST("/", controller.LogIn)
	r.GET("/signup", controller.SignUpPage)
	r.POST("/signup", controller.SignUp)
	r.GET("/logout", controller.LogOut)
	r.GET("/dashboard", controller.DashBoardPage)

	r.GET("/lists", controller.AllLists)
	r.POST("/lists", controller.CreateList)
	r.GET("/lists/:user", controller.FindListByUserName)
	r.POST("/lists/delete/:id", controller.DeleteListById)

	r.Run()
}
