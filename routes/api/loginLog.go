package api

import (
	"mugg/guapin/app/controllers"

	"github.com/gin-gonic/gin"
)

// LoginLog is
func LoginLog(g *gin.RouterGroup) {
	NewLoginLogController := controllers.NewLoginLog()

	storage := g.Group("/login_log")
	// storage.Use(jwtauth.JWTAuth())
	{
		storage.GET("", NewLoginLogController.Home)
		storage.POST("", NewLoginLogController.Create)
		storage.PUT("", NewLoginLogController.Update)
		storage.DELETE("", NewLoginLogController.Delete)
	}
}
