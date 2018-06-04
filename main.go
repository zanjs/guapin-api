package main

import (
	"mugg/guapin/app/conf"
	"mugg/guapin/app/middleware"
	"mugg/guapin/routes"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.Use(cors.Default())
	r.Use(middleware.Options)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	config := conf.Config
	r.LoadHTMLGlob("templates/**/*")
	routes.Web(r)
	routes.API(r)

	r.Run(":" + config.APP.Port)
}
