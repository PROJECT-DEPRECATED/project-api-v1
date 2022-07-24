package routes

import (
	"github.com/gin-gonic/gin"
)

func Resources(app *gin.Engine) {
	rsrc := app.Group("/resources")
	{
		rsrc.StaticFile("/favicon", "./resources/favicon.ico")
		rsrc.StaticFile("/logo", "./resources/logo.png")
	}
}
