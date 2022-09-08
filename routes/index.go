package routes

import (
	"github.com/gin-gonic/gin"
)

func Index(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(301, "https://github.com/devproje/project-api")
	})
}
