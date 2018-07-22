package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Options is
func Options(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE")
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	c.Header("Allow", "*")
	c.Header("Content-Type", "application/json")
	// c.Header("Content-Type", "charset=utf-8, application/json")
	if c.Request.Method != "OPTIONS" {
		fmt.Println(c.Request.RequestURI)
		fmt.Println("!===OPTIONS")
		c.Next()
	} else {
		fmt.Println(c.Request.RequestURI)
		fmt.Println("=OPTIONS")
		c.AbortWithStatus(http.StatusOK)
		c.Abort()
	}
}
