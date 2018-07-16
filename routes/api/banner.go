package api

import (
	"mugg/guapin/app/controllers"

	"github.com/gin-gonic/gin"
)

// Banner is
func Banner(g *gin.RouterGroup) {
	NewBannerController := controllers.NewBanner()

	Banner := g.Group("/banner")
	// Banner.Use(jwtauth.JWTAuth())
	{
		Banner.GET("", NewBannerController.Home)
		Banner.POST("", NewBannerController.Create)
		Banner.PUT("", NewBannerController.Update)
		Banner.DELETE("", NewBannerController.Delete)
	}
}
