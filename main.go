package main

import (
	"io"
	"mugg/guapin/app/conf"
	"mugg/guapin/app/middleware"
	"mugg/guapin/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/upload", "./upload")
	f, _ := os.Create("g.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(cors.Default())
	// r.LoadHTMLGlob("templates/**/*")
	r.Use(middleware.Options)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	config := conf.Config
	routes.Web(r)
	routes.API(r)

	r.Run(":" + config.APP.Port)
}
