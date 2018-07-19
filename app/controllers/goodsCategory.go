package controllers

import (
	"mugg/guapin/app/service"
	"mugg/guapin/model"

	"github.com/gin-gonic/gin"
)

type (
	// GoodsCategoryController is
	GoodsCategoryController struct {
		BaseController
		GoodsCategory service.GoodsCategory
	}
)

// NewGoodsCategory is
func NewGoodsCategory() *GoodsCategoryController {
	return &GoodsCategoryController{}
}

// Home is
func (s GoodsCategoryController) Home(c *gin.Context) {

	disabled := c.DefaultQuery("disabled", "")

	maps := make(map[string]interface{})
	// categoryID := searchData.CategoryID
	if disabled != "" {
		maps["disabled"] = disabled
	}

	data, err := s.GoodsCategory.GetAllQuerySearch(maps)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	// s.SuccessJSONDataPage(c, count, data)
	s.SuccessJSONData(c, data)
}

// Create is
func (s GoodsCategoryController) Create(c *gin.Context) {
	GoodsCategory := &model.GoodsCategory{}

	if err := c.BindJSON(GoodsCategory); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	if GoodsCategory.Name == "" {
		s.ErrorJSON(c, "name is null")
		return
	}

	err := s.GoodsCategory.Create(GoodsCategory)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, GoodsCategory)
}

// Update is
func (s GoodsCategoryController) Update(c *gin.Context) {
	category := &model.GoodsCategory{}

	if err := c.BindJSON(category); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if category.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data, err := s.GoodsCategory.Get(category.ID)
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
func (s GoodsCategoryController) Delete(c *gin.Context) {
	category := &model.GoodsCategory{}

	if err := c.BindJSON(category); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if category.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data, err := s.GoodsCategory.Get(category.ID)
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
