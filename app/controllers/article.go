package controllers

import (
	"fmt"
	"mugg/guapin/app/service"
	"mugg/guapin/model"
	// "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type (
	// ArticleController is
	ArticleController struct {
		BaseController
		Article  service.Article
		Category service.Category
	}
)

// NewArticle is
func NewArticle() *ArticleController {
	return &ArticleController{}
}

// Home is
func (s ArticleController) Home(c *gin.Context) {
	data, err := s.Article.GetAll()
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, data)
}

// Create is
func (s ArticleController) Create(c *gin.Context) {
	data := &model.Article{}

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

	err = s.Article.Create(data)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, data)
}

// Update is
func (s ArticleController) Update(c *gin.Context) {
	data := &model.Article{}

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

	data2, err := s.Article.Get(data.ID)
	data2.Title = data.Title
	data2.Description = data.Description
	data2.CategoryID = data.CategoryID
	data2.Content = data.Content
	data2.Update(data)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONUpdate(c)
}

// Delete is
func (s ArticleController) Delete(c *gin.Context) {
	category := &model.Article{}

	if err := c.BindJSON(category); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if category.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data, err := s.Article.Get(category.ID)
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
