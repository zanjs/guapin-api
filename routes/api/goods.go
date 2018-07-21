package api

import (
	"mugg/guapin/app/controllers"

	"github.com/gin-gonic/gin"
)

// Goods is
func Goods(g *gin.RouterGroup) {
	NewGoodsController := controllers.NewGoods()
	NewGoodsPictureController := controllers.NewGoodsPicture()

	Goods := g.Group("/goods")
	// Goods.Use(jwtauth.JWTAuth())
	{
		Goods.GET("", NewGoodsController.Home)
		Goods.GET("/:id", NewGoodsController.Get)
		Goods.POST("", NewGoodsController.Create)
		Goods.PUT("", NewGoodsController.Update)
		Goods.DELETE("", NewGoodsController.Delete)
	}

	GoodsPicture := g.Group("/goods_picture")
	{
		GoodsPicture.POST("", NewGoodsPictureController.Create)
		GoodsPicture.DELETE("", NewGoodsPictureController.Delete)
	}
}
