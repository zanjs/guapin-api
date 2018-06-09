package api

import (
	"mugg/guapin/app/controllers"

	"github.com/gin-gonic/gin"
)

// Logs is
func Logs(g *gin.RouterGroup) {
	NewLogsController := controllers.NewLogs()

	Logs := g.Group("/logs")
	{
		Logs.POST("", NewLogsController.Create)
	}
}
