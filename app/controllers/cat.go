package controllers

import (
	"github.com/gin-gonic/gin"
)

type (
	// CatController is
	CatController struct{}
)

// NewCat is
func NewCat() *CatController {
	return &CatController{}
}

// Home is
func (s CatController) Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"va": "123",
	})
}
