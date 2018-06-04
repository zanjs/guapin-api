package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Web si
func Web(g *gin.Engine) {
	g.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "hello zan!",
		})
	})

	g.StaticFS("/upload", http.Dir("upload"))
}
