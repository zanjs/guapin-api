package response

import (
	"github.com/gin-gonic/gin"
)

// SuccessJSON is ok
func SuccessJSON(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
}

// ErrorJSON is ok
func ErrorJSON(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"success": false,
		"error":   msg,
	})
}

// SuccessJSONData is ok
func SuccessJSONData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}

// SuccessJSONDataPage is ok
func SuccessJSONDataPage(c *gin.Context, total int, data interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"total":   total,
		"data":    data,
	})
}

// SuccessJSONToken is ok
func SuccessJSONToken(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"token":   data,
	})
}

// SuccessJSONUpdate is ok
func SuccessJSONUpdate(c *gin.Context) {
	SuccessJSONData(c, "更新成功")
}

// SuccessJSONDelete is
func SuccessJSONDelete(c *gin.Context) {
	SuccessJSONData(c, "删除成功")
}
