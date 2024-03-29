package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/devproje/plog/log"
	"github.com/devproje/project-website/config"
	"github.com/devproje/project-website/middleware"
	"github.com/devproje/project-website/routes"
	"github.com/devproje/project-website/utils"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

const (
	VERSION = "v1.1.0-beta.1"
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
		err := os.WriteFile("config.json", []byte(config.GetSample), 0755)
		if err != nil {
			log.Fatalf("failed to create 'config.json'\n%v", err)
		}

		log.Fatalf("'config.json' isn't exist!\n%v", err)
	}

	_, err = os.ReadDir("./.file")
	if err != nil {
		err := os.Mkdir("./.file", 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	err = utils.Connect()
	if err != nil {
		log.Errorln(err)
	}

	app := gin.Default()

	app.Use(favicon.New("./resources/favicon.ico"))
	app.Use(middleware.Cors)

	routes.Router(app)

	app.Run(fmt.Sprintf(":%d", port))
}
