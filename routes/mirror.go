package routes

import (
	"github.com/gin-gonic/gin"
)

func Mirror(app *gin.Engine) {
	group := app.Group("mirror")
	{
		group.Static("/file", "./file")
	}
}
