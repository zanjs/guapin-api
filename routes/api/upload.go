package api

import (
	"mugg/guapin/app/controllers"

	"github.com/gin-gonic/gin"
)

// Upload is
func Upload(g *gin.RouterGroup) {
	NewUploadController := controllers.NewUpload()

	Upload := g.Group("/upload")
	// Upload.Use(jwtauth.JWTAuth())
	{
		Upload.POST("", NewUploadController.Create)
		Upload.DELETE("", NewUploadController.Delete)
		Upload.GET("/config", NewUploadController.ConfigInfo)
	}
}
