package api

import (
	"mugg/guapin/app/controllers"
	"mugg/guapin/app/middleware/jwtauth"

	"github.com/gin-gonic/gin"
)

// AdminUser is
func AdminUser(g *gin.RouterGroup) {
	NewAdminUserController := controllers.NewAdminUser()
	g.GET("/my", NewAdminUserController.GetMe)
	user := g.Group("/user")
	user.Use(jwtauth.JWTAuth())
	{
		user.GET("", NewAdminUserController.Home)
		user.GET("/my", NewAdminUserController.GetMe)
		user.POST("", NewAdminUserController.Create)
		user.PUT("", NewAdminUserController.Update)
	}

	login := g.Group("/login")
	{
		login.POST("", NewAdminUserController.Login)
		// login.POST("", middleware.AuthMiddleware.LoginHandler)
	}
}
