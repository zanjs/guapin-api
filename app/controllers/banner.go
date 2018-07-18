package controllers

import (
	"fmt"
	"mugg/guapin/app/conf"
	"mugg/guapin/app/middleware"
	"mugg/guapin/app/service"
	"mugg/guapin/model"
	"strconv"
	"strings"
	// "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type (
	// BannerController is
	BannerController struct {
		BaseController
		Banner   service.Banner
		Category service.BannerCategory
	}
)

// NewBanner is
func NewBanner() *BannerController {
	return &BannerController{}
}

// Home is
func (s BannerController) Home(c *gin.Context) {

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

	// data, err := s.Banner.GetAll()
	data, err := s.Banner.GetAllQuerySearch(qPage, maps, searchData.Name)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	count, err := s.Banner.GetAllQuerySearchTotal(maps, searchData.Name)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	fileInfo := conf.Config.File

	fmt.Println(fileInfo)
	for i := range data {
		// fmt.Println(k, v)

		if data[i].Picture != "" {

			picture := data[i].Picture
			picture4 := picture[0:4]
			fmt.Println(picture4)
			if picture4 != "http" {
				data[i].Picture = fileInfo.Host + data[i].Picture
			}
			// data[i].Content = ""
		}
	}

	s.SuccessJSONDataPage(c, count, data)
}

// Get is
func (s BannerController) Get(c *gin.Context) {
	id := c.Param("id")
	uID, _ := strconv.ParseUint(id, 0, 64)
	data, err := s.Banner.Get(uID)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	s.SuccessJSONData(c, data)
}

// Create is
func (s BannerController) Create(c *gin.Context) {
	data := &model.Banner{}

	if err := c.BindJSON(data); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if data.CategoryID == 0 {
		s.ErrorJSON(c, "栏目不能为空")
		return
	}

	// categoryID, _ := strconv.ParseUint(data.CategoryID, 10, 64)
	categoryID := data.CategoryID

	fmt.Println(categoryID)

	category, err := s.Category.Get(categoryID)
	if category.Name == "" {
		s.ErrorJSON(c, "没有找到栏目")
		return
	}

	cdata := &model.Banner{}
	data.Picture = strings.Replace(data.Picture, conf.Config.File.Host, "", -1)
	cdata.Name = data.Name
	cdata.Description = data.Description
	cdata.CategoryID = data.CategoryID
	cdata.Picture = data.Picture
	err = s.Banner.Create(cdata)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, cdata)
}

// Update is
func (s BannerController) Update(c *gin.Context) {
	data := &model.Banner{}

	if err := c.BindJSON(data); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if data.CategoryID == 0 {
		s.ErrorJSON(c, "类别不能为空")
		return
	}

	// categoryID, _ := strconv.ParseUint(data.CategoryID, 10, 64)
	categoryID := data.CategoryID

	fmt.Println(categoryID)

	category, err := s.Category.Get(categoryID)
	if category.Name == "" {
		s.ErrorJSON(c, "没有找到类别")
		return
	}

	data.Picture = strings.Replace(data.Picture, conf.Config.File.Host, "", -1)

	data2, err := s.Banner.GetID(data.ID)
	data2.Name = data.Name
	data2.Description = data.Description
	data2.CategoryID = data.CategoryID
	data2.Picture = data.Picture
	data2.Sort = data.Sort
	data2.Update(&data2)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONUpdate(c)
}

// Delete is
func (s BannerController) Delete(c *gin.Context) {
	category := &model.Banner{}

	if err := c.BindJSON(category); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if category.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data, err := s.Banner.Get(category.ID)
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
