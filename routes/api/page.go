package api

import (
	"mugg/guapin/app/controllers"

	"github.com/gin-gonic/gin"
)

// Page is
func Page(g *gin.RouterGroup) {
	NewPageController := controllers.NewPage()

	Page := g.Group("/page")
	// Page.Use(jwtauth.JWTAuth())
	{
		Page.GET("", NewPageController.Home)
		Page.POST("", NewPageController.Create)
		Page.PUT("", NewPageController.Update)
		Page.DELETE("", NewPageController.Delete)
	}
}
