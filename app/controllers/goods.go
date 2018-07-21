package controllers

import (
	"fmt"
	"mugg/guapin/app/middleware"
	"mugg/guapin/app/service"
	"mugg/guapin/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	// GoodsController is
	GoodsController struct {
		BaseController
		Goods    service.Goods
		Category service.GoodsCategory
	}
)

// NewGoods is
func NewGoods() *GoodsController {
	return &GoodsController{}
}

// Home is
func (s GoodsController) Home(c *gin.Context) {

	qPage := middleware.GetPage(c)

	fmt.Println("qPage\b")
	fmt.Println(qPage)

	searchData := model.SearchQ{}
	searchData.Name = c.DefaultQuery("name", "")
	categoryID := c.DefaultQuery("category_id", "")

	fmt.Println("searchData")
	fmt.Println(searchData)

	maps := make(map[string]interface{})
	// categoryID := searchData.CategoryID
	if categoryID != "" {
		maps["category_id"] = categoryID
	}

	// data, err := s.Goods.GetAll()
	data, err := s.Goods.GetAllQuerySearch(qPage, maps, searchData.Name)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	count, err := s.Goods.GetAllQuerySearchTotal(maps, searchData.Name)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	s.SuccessJSONDataPage(c, count, data)
}

// Get is
func (s GoodsController) Get(c *gin.Context) {
	id := c.Param("id")
	uID, _ := strconv.ParseUint(id, 0, 64)
	data, err := s.Goods.Get(uID)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	s.SuccessJSONData(c, data)
}

// Create is
func (s GoodsController) Create(c *gin.Context) {
	data := &model.Goods{}

	if err := c.BindJSON(data); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if data.CategoryID == 0 {
		s.ErrorJSON(c, "分类不能为空")
		return
	}

	// categoryID, _ := strconv.ParseUint(data.CategoryID, 10, 64)
	categoryID := data.CategoryID

	fmt.Println(categoryID)

	category, err := s.Category.Get(categoryID)
	if category.Name == "" {
		s.ErrorJSON(c, "没有找到分类")
		return
	}

	cdata := &model.Goods{}
	// data.Picture = strings.Replace(data.Picture, conf.Config.File.Host, "", -1)
	cdata.Name = data.Name
	cdata.Description = data.Description
	cdata.CategoryID = data.CategoryID
	// cdata.Picture = data.Picture
	err = s.Goods.Create(cdata)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, cdata)
}

// Update is
func (s GoodsController) Update(c *gin.Context) {
	data := &model.Goods{}

	if err := c.BindJSON(data); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if data.CategoryID == 0 {
		s.ErrorJSON(c, "分类不能为空")
		return
	}

	// categoryID, _ := strconv.ParseUint(data.CategoryID, 10, 64)
	categoryID := data.CategoryID

	fmt.Println(categoryID)

	category, err := s.Category.Get(categoryID)
	if category.Name == "" {
		s.ErrorJSON(c, "没有找到分类")
		return
	}

	// data.Picture = strings.Replace(data.Picture, conf.Config.File.Host, "", -1)

	data2, err := s.Goods.GetID(data.ID)
	data2.Name = data.Name
	data2.Description = data.Description
	data2.CategoryID = data.CategoryID
	data2.Sort = data.Sort
	data2.RecommendStatus = data.RecommendStatus
	data2.Status = data.Status
	data2.Stores = data.Stores
	data2.PinTuanPrice = data.PinTuanPrice
	data2.MinPrice = data.MinPrice
	data2.OriginalPrice = data.OriginalPrice
	data2.MinScore = data.MinScore
	data2.PingTuan = data.PingTuan
	data2.Weight = data.Weight
	data2.Picture = data.Picture
	data2.Update(&data2)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONUpdate(c)
}

// Delete is
func (s GoodsController) Delete(c *gin.Context) {
	category := &model.Goods{}

	if err := c.BindJSON(category); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if category.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data, err := s.Goods.Get(category.ID)
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
