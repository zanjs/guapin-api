package api

import (
	"mugg/guapin/app/controllers"

	"github.com/gin-gonic/gin"
)

// GoodsCategory is
func GoodsCategory(g *gin.RouterGroup) {
	NewGoodsCategoryController := controllers.NewGoodsCategory()

	GoodsCategory := g.Group("/goods_category")
	// GoodsCategory.Use(jwtauth.JWTAuth())
	{
		GoodsCategory.GET("", NewGoodsCategoryController.Home)
		GoodsCategory.POST("", NewGoodsCategoryController.Create)
		GoodsCategory.PUT("", NewGoodsCategoryController.Update)
		GoodsCategory.DELETE("", NewGoodsCategoryController.Delete)
	}
}
