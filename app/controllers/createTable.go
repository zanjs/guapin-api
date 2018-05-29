package controllers

import (
	"mugg/guapin/model"

	"github.com/gin-gonic/gin"
)

type (
	// CreateTableController is
	CreateTableController struct {
		BaseController
	}
)

// NewCreateTable is
func NewCreateTable() *CreateTableController {
	return &CreateTableController{}
}

// Home is
func (s CreateTableController) Home(c *gin.Context) {
	err := model.CreateTable()

	if err != nil {
		s.ErrorJSON(c, err.Error())
	}
	s.SuccessJSON(c)
}
