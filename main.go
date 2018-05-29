package main

import (
	"mugg/guapin/app/conf"
	"mugg/guapin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := conf.Config
	r.LoadHTMLGlob("templates/**/*")
	routes.Web(r)
	routes.API(r)

	r.Run(":" + config.APP.Port)
}
