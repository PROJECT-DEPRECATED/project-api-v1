package routes

import (
	"github.com/devproje/project-website/config"
	"github.com/gin-gonic/gin"
)

func Index(app *gin.Engine) {
	conf, _ := config.Get()

	app.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(301, conf.URL+"/api")
	})
}
