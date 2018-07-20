package controllers

import (
	"fmt"
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

func buildData(list []*model.GoodsCategory) map[uint64]map[uint64]*model.GoodsCategory {
	var data map[uint64]map[uint64]*model.GoodsCategory = make(map[uint64]map[uint64]*model.GoodsCategory)
	for _, v := range list {
		id := v.ID
		fid := v.PID
		if _, ok := data[fid]; !ok {
			data[fid] = make(map[uint64]*model.GoodsCategory)
		}
		data[fid][id] = v
	}
	return data
}

func makeTreeCore(index uint64, data map[uint64]map[uint64]*model.GoodsCategory) []*model.GoodsCategory {
	tmp := make([]*model.GoodsCategory, 0)
	for id, item := range data[index] {
		if data[id] != nil {
			item.List = makeTreeCore(id, data)
		}
		tmp = append(tmp, item)
	}
	return tmp
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

	tmp := make([]model.GoodsCategory, 0)

	for _, v := range data {
		id := v.ID
		pid := v.PID
		fmt.Println(id)
		fmt.Println(v)
		if pid == 0 {
			tmp = append(tmp, v)
			for _, v2 := range data {
				if v2.PID == v.ID {
					tmp = append(tmp, v2)
				}
			}
		}
	}

	fmt.Println(tmp)

	// s.SuccessJSONDataPage(c, count, data)
	s.SuccessJSONData(c, tmp)
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

	if GoodsCategory.PID > 0 {
		data, err := s.GoodsCategory.Get(GoodsCategory.PID)
		if err != nil {
			s.ErrorJSON(c, err.Error())
			return
		}
		GoodsCategory.Level = (data.Level + 1)
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
