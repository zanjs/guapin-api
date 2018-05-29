package api

import (
	"mugg/guapin/app/controllers"

	"github.com/gin-gonic/gin"
)

// Category is
func Category(g *gin.RouterGroup) {
	NewCategoryController := controllers.NewCategory()

	Category := g.Group("/category")
	// Category.Use(jwtauth.JWTAuth())
	{
		Category.GET("", NewCategoryController.Home)
		Category.POST("", NewCategoryController.Create)
		Category.PUT("", NewCategoryController.Update)
		Category.DELETE("", NewCategoryController.Delete)
	}
}
