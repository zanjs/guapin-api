package controllers

import (
	"mugg/guapin/app/service"
	"mugg/guapin/model"

	"github.com/gin-gonic/gin"
)

type (
	// BannerCategoryController is
	BannerCategoryController struct {
		BaseController
		BannerCategory service.BannerCategory
	}
)

// NewBannerCategory is
func NewBannerCategory() *BannerCategoryController {
	return &BannerCategoryController{}
}

// Home is
func (s BannerCategoryController) Home(c *gin.Context) {
	data, err := s.BannerCategory.GetAll()
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	// s.SuccessJSONDataPage(c, count, data)
	s.SuccessJSONData(c, data)
}

// Create is
func (s BannerCategoryController) Create(c *gin.Context) {
	BannerCategory := &model.BannerCategory{}

	if err := c.BindJSON(BannerCategory); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	if BannerCategory.Name == "" {
		s.ErrorJSON(c, "name is null")
		return
	}

	err := s.BannerCategory.Create(BannerCategory)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, BannerCategory)
}

// Update is
func (s BannerCategoryController) Update(c *gin.Context) {
	category := &model.BannerCategory{}

	if err := c.BindJSON(category); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if category.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data, err := s.BannerCategory.Get(category.ID)
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
func (s BannerCategoryController) Delete(c *gin.Context) {
	category := &model.BannerCategory{}

	if err := c.BindJSON(category); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if category.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data, err := s.BannerCategory.Get(category.ID)
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
