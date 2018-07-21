package controllers

import (
	"mugg/guapin/app/service"
	"mugg/guapin/model"

	"github.com/gin-gonic/gin"
)

type (
	// GoodsPictureController is
	GoodsPictureController struct {
		BaseController
		GoodsPicture service.GoodsPicture
		Goods        service.Goods
	}
)

// NewGoodsPicture is
func NewGoodsPicture() *GoodsPictureController {
	return &GoodsPictureController{}
}

// Home is
func (s GoodsPictureController) Home(c *gin.Context) {

	disabled := c.DefaultQuery("disabled", "")

	maps := make(map[string]interface{})
	// categoryID := searchData.CategoryID
	if disabled != "" {
		maps["disabled"] = disabled
	}

	data, err := s.GoodsPicture.GetAllQuerySearch(maps)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	s.SuccessJSONData(c, data)
}

// Create is
func (s GoodsPictureController) Create(c *gin.Context) {
	GoodsPicture := &model.GoodsPicture{}

	if err := c.BindJSON(GoodsPicture); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	if GoodsPicture.Picture == "" {
		s.ErrorJSON(c, "Picture is null")
		return
	}

	if GoodsPicture.GoodsID == 0 {
		s.ErrorJSON(c, "GoodsID is null")
		return
	}

	_, err := s.Goods.GetID(GoodsPicture.GoodsID)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	err = s.GoodsPicture.Create(GoodsPicture)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, GoodsPicture)
}

// Update is
func (s GoodsPictureController) Update(c *gin.Context) {
	category := &model.GoodsPicture{}

	if err := c.BindJSON(category); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if category.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data, err := s.GoodsPicture.Get(category.ID)
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
func (s GoodsPictureController) Delete(c *gin.Context) {
	data := &model.GoodsPicture{}

	if err := c.BindJSON(data); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if data.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data2, err := s.GoodsPicture.Get(data.ID)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	err = data2.Delete()
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONUpdate(c)
}
