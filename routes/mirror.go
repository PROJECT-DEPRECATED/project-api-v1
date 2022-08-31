package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

func quaryDir() []string {
	list := []string{}
	filepath := "./file"
	dir, _ := os.ReadDir(filepath)
	for _, j := range dir {
		list = append(list, j.Name())
	}

	return list
}

func Mirror(app *gin.Engine) {
	group := app.Group("mirror")
	{
		group.Static("/file", "./file")
	}
}
