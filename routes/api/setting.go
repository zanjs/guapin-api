package api

import (
	"mugg/guapin/app/controllers"
	"mugg/guapin/app/middleware/jwtauth"

	"github.com/gin-gonic/gin"
)

// Setting is
func Setting(g *gin.RouterGroup) {
	NewSettingController := controllers.NewSetting()

	Setting := g.Group("/setting")
	Setting.Use(jwtauth.JWTAuth())
	{
		Setting.GET("", NewSettingController.Home)
		Setting.POST("", NewSettingController.Create)
		Setting.PUT("", NewSettingController.Update)
		Setting.DELETE("", NewSettingController.Delete)
	}
}
