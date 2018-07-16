package api

import (
	"mugg/guapin/app/controllers"

	"github.com/gin-gonic/gin"
)

// BannerCategory is
func BannerCategory(g *gin.RouterGroup) {
	NewBannerCategoryController := controllers.NewBannerCategory()

	BannerCategory := g.Group("/banner_category")
	// BannerCategory.Use(jwtauth.JWTAuth())
	{
		BannerCategory.GET("", NewBannerCategoryController.Home)
		BannerCategory.POST("", NewBannerCategoryController.Create)
		BannerCategory.PUT("", NewBannerCategoryController.Update)
		BannerCategory.DELETE("", NewBannerCategoryController.Delete)
	}
}
