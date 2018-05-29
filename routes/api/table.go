package api

import (
	"mugg/guapin/app/controllers"

	"github.com/gin-gonic/gin"
)

// Table is
func Table(g *gin.RouterGroup) {
	NewCreateTableController := controllers.NewCreateTable()

	r := g.Group("/table")
	{
		r.GET("", NewCreateTableController.Home)
	}
}
