package controllers

import (
	"mugg/guapin/app/service"
	"mugg/guapin/model"

	"github.com/gin-gonic/gin"
)

type (
	// CategoryController is
	CategoryController struct {
		BaseController
		Category service.Category
	}
)

// NewCategory is
func NewCategory() *CategoryController {
	return &CategoryController{}
}

// Home is
func (s CategoryController) Home(c *gin.Context) {
	data, err := s.Category.GetAll()
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, data)
}

// Create is
func (s CategoryController) Create(c *gin.Context) {
	Category := &model.Category{}

	if err := c.BindJSON(Category); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	err := s.Category.Create(Category)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, Category)
}

// Update is
func (s CategoryController) Update(c *gin.Context) {
	category := &model.Category{}

	if err := c.BindJSON(category); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if category.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data, err := s.Category.Get(category.ID)
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
func (s CategoryController) Delete(c *gin.Context) {
	category := &model.Category{}

	if err := c.BindJSON(category); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if category.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data, err := s.Category.Get(category.ID)
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
