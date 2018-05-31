package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Options is
func Options(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	c.Header("Allow", "*")
	// c.Header("Content-Type", "*/*")
	c.Header("Content-Type", "charset=utf-8, application/json")
	if c.Request.Method != "OPTIONS" {
		fmt.Println("OPTIONS")
		c.Next()
	} else {
		// c.Header("Access-Control-Allow-Origin", "*")
		// c.Header("Access-Control-Allow-Methods", "*")
		// c.Header("Access-Control-Allow-Headers", "*")
		// c.Header("Allow", "*")
		// c.Header("Content-Type", "*")
		fmt.Println("opppsa")
		c.AbortWithStatus(http.StatusOK)
		c.Abort()
	}
}
