package api

import (
	"mugg/guapin/app/controllers"

	"github.com/gin-gonic/gin"
)

// Storage is
func Storage(g *gin.RouterGroup) {
	NewStorageController := controllers.NewStorage()

	storage := g.Group("/storage")
	// storage.Use(jwtauth.JWTAuth())
	{
		storage.GET("", NewStorageController.Home)
		storage.POST("", NewStorageController.Create)
		storage.PUT("", NewStorageController.Update)
		storage.DELETE("", NewStorageController.Delete)
	}
}
