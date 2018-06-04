package api

import (
	"mugg/guapin/app/controllers"

	"github.com/gin-gonic/gin"
)

// Article is
func Article(g *gin.RouterGroup) {
	NewArticleController := controllers.NewArticle()

	Article := g.Group("/article")
	// Article.Use(jwtauth.JWTAuth())
	{
		Article.GET("", NewArticleController.Home)
		Article.GET("/:id", NewArticleController.Get)
		Article.POST("", NewArticleController.Create)
		Article.PUT("", NewArticleController.Update)
		Article.DELETE("", NewArticleController.Delete)
	}
}
