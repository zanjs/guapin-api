package api

import (
	"mugg/guapin/app/controllers"
	"mugg/guapin/app/middleware/jwtauth"

	"github.com/gin-gonic/gin"
)

// User is
func User(g *gin.RouterGroup) {
	NewUserController := controllers.NewUser()

	user := g.Group("/user")
	user.Use(jwtauth.JWTAuth())
	{
		user.GET("", NewUserController.Home)
		// user.GET("/my", jwtauth.JWTAuth())
		user.POST("", NewUserController.Create)
		user.PUT("", NewUserController.Update)
	}

	login := g.Group("/login")
	{
		login.POST("", NewUserController.Login)
		// login.POST("", middleware.AuthMiddleware.LoginHandler)
	}
}
