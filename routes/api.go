package routes

import (
	"github.com/devproje/project-website/api"
	"github.com/gin-gonic/gin"
)

func APIV1(app *gin.Engine) {
	v1 := app.Group("/v1")
	{
		v1.GET("/current", api.CurrentTime)
		v1.GET("/hangang", api.Hangang)
		v1.GET("/hangang/:area", api.Hangang)
		v1.GET("/mcprofile/:username", api.MCProfile)
		v1.GET("/led", api.GetLed)
		v1.POST("/led", api.SetLed)

		// blog := app.Group("/blog")
		// {
		// }
	}
}
