package main

import (
	"flag"
	"fmt"

	"github.com/devproje/project-website/api"
	"github.com/devproje/project-website/config"
	"github.com/devproje/project-website/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

var (
	port  int
	debug bool
)

func init() {
	flag.IntVar(&port, "port", 3000, "Service port")
	flag.BoolVar(&debug, "debug", false, "Debug mode")
	flag.Parse()
}

func main() {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	utils.Connect()

	app := gin.Default()
	cor := cors.DefaultConfig()

	app.Use(favicon.New("./frontend/public/favicon.ico"))

	cor.AllowOrigins = []string{"*"}
	app.Use(cors.New(cor))

	app.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(301, config.Get().URL+"/api")
	})

	v1 := app.Group("/v1")
	{
		v1.GET("/hangang", api.Hangang)
		v1.GET("/hangang/:area", api.Hangang)

		v1.GET("/current", api.CurrentTime)
	}

	app.Run(fmt.Sprintf(":%d", port))
}
