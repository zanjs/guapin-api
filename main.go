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

	config := conf.Config
	r.LoadHTMLGlob("templates/**/*")
	routes.Web(r)
	routes.API(r)

	r.Run(":" + config.APP.Port)
}
