package api

import (
	"mugg/guapin/app/controllers"

	"github.com/gin-gonic/gin"
)

// Excel is
func Excel(g *gin.RouterGroup) {
	NewExcelController := controllers.NewExcel()

	Excel := g.Group("/excel")
	// Excel.Use(jwtauth.JWTAuth())
	{
		Excel.POST("", NewExcelController.Create)
	}
}
