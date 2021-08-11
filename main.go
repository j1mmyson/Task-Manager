package main

import (
	"embed"
	"html/template"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/j1mmyson/reviewList/api"
	"github.com/j1mmyson/reviewList/controller"
	"github.com/j1mmyson/reviewList/models"
)

var (
	//go:embed web/templates/*
	templatesFS embed.FS
	//go:embed web
	staticFS embed.FS
)

func main() {

	r := gin.Default()
	LoadHTMLFromEmbedFS(r, templatesFS, "web/templates/*")

	r.GET("/static/*filepath", func(c *gin.Context) {
		c.FileFromFS(path.Join("/web/", c.Request.URL.Path), http.FS(staticFS))
	})

	models.ConnectDB()

	r.GET("/", controller.LogInPage)
	r.POST("/", controller.LogIn)
	r.GET("/signup", controller.SignUpPage)
	r.POST("/signup", controller.SignUp)
	r.GET("/logout", controller.LogOut)
	r.GET("/dashboard", controller.DashBoardPage)

	r.GET("/lists", controller.AllLists)
	r.POST("/lists", controller.CreateList)
	r.GET("/lists/:user", controller.FindListByUserName)
	r.POST("/delete/:id", controller.DeleteListById)
	r.POST("/edit/:id", controller.EditListById)

	apiRouter := r.Group("/api")
	{
		apiRouter.POST("/user", api.CreateUser)
		apiRouter.GET("/user", api.ShowUserList)
		apiRouter.GET("/user/:id", api.GetUser)
		apiRouter.DELETE("/user", api.DeleteUser)

		apiRouter.GET("/:user_id/card", api.GetCards)
		apiRouter.POST("/card", api.CreateCard)
		apiRouter.DELETE("/card/:id", api.DeleteCard)
		apiRouter.PUT("/card/:id", api.EditCard)
	}

	r.Run(":8080")
}

func LoadHTMLFromEmbedFS(r *gin.Engine, em embed.FS, pattern string) {
	templ := template.Must(template.ParseFS(em, pattern))

	r.SetHTMLTemplate(templ)
}
