package controllers

import (
	"mugg/guapin/app/response"

	"github.com/gin-gonic/gin"
)

// BaseController is
type BaseController struct{}

// SuccessJSON is ok
func (b BaseController) SuccessJSON(c *gin.Context) {
	response.SuccessJSON(c)
}

// ErrorJSON is ok
func (b BaseController) ErrorJSON(c *gin.Context, msg string) {
	response.ErrorJSON(c, msg)
}

// SuccessJSONData is ok
func (b BaseController) SuccessJSONData(c *gin.Context, data interface{}) {
	response.SuccessJSONData(c, data)
}

// SuccessJSONDataPage is
func (b BaseController) SuccessJSONDataPage(c *gin.Context, total int, data interface{}) {
	response.SuccessJSONDataPage(c, total, data)
}

// SuccessJSONUpdate is ok
func (b BaseController) SuccessJSONUpdate(c *gin.Context) {
	response.SuccessJSONUpdate(c)
}

// SuccessJSONDelete is ok
func (b BaseController) SuccessJSONDelete(c *gin.Context) {
	response.SuccessJSONDelete(c)
}
