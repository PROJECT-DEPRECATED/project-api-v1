package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/devproje/project-website/config"
	"github.com/devproje/project-website/log"
	"github.com/devproje/project-website/routes"
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
	_, err := config.Get()
	if err != nil {
		err := os.WriteFile("config.json", []byte(config.GetSample), 0666)
		if err != nil {
			log.Logger.Fatalf("failed to create 'config.json'\n%v", err)
		}

		log.Logger.Fatalf("'config.json' isn't exist!\n%v", err)
	}

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	err = utils.Connect()
	if err != nil {
		log.Logger.Errorln(err)
	}

	app := gin.Default()
	cor := cors.DefaultConfig()

	app.Use(favicon.New("./resources/favicon.ico"))

	cor.AllowOrigins = []string{"*"}
	app.Use(cors.New(cor))

	routes.Index(app)
	routes.APIV1(app)
	routes.Resources(app)

	app.Run(fmt.Sprintf(":%d", port))
}
