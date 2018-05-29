package controllers

import (
	"mugg/guapin/app/service"
	"mugg/guapin/model"

	"github.com/gin-gonic/gin"
)

type (
	// PageController is
	PageController struct {
		BaseController
		Page service.Page
	}
)

// NewPage is
func NewPage() *PageController {
	return &PageController{}
}

// Home is
func (s PageController) Home(c *gin.Context) {
	data, err := s.Page.GetAll()
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, data)
}

// Create is
func (s PageController) Create(c *gin.Context) {
	Page := &model.Page{}

	if err := c.BindJSON(Page); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	err := s.Page.Create(Page)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, Page)
}

// Update is
func (s PageController) Update(c *gin.Context) {
	category := &model.Page{}

	if err := c.BindJSON(category); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if category.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data, err := s.Page.Get(category.ID)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	err = data.Update(category)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONUpdate(c)
}

// Delete is
func (s PageController) Delete(c *gin.Context) {
	category := &model.Page{}

	if err := c.BindJSON(category); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if category.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data, err := s.Page.Get(category.ID)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	err = data.Delete()
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONUpdate(c)
}
