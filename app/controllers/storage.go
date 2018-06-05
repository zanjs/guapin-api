package controllers

import (
	"fmt"
	"mugg/guapin/app/conf"
	"mugg/guapin/app/middleware"
	"mugg/guapin/app/service"
	"mugg/guapin/model"

	"github.com/gin-gonic/gin"
)

type (
	// StorageController is
	StorageController struct {
		BaseController
		Storage service.Storage
	}
)

// NewStorage is
func NewStorage() *StorageController {
	return &StorageController{}
}

// Home is
func (s StorageController) Home(c *gin.Context) {
	qPage := middleware.GetPage(c)

	// data, err := s.Storage.GetAllQuery(qPage)
	// if err != nil {
	// 	s.ErrorJSON(c, err.Error())
	// 	return
	// }
	// count, err := s.Storage.GetAllQueryTotal()
	// if err != nil {
	// 	s.ErrorJSON(c, err.Error())
	// 	return
	// }

	searchName := c.DefaultQuery("title", "")

	data, err := s.Storage.GetAllQuerySearch(qPage, searchName)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	count, err := s.Storage.GetAllQuerySearchTotal(searchName)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	fileInfo := conf.Config.File

	fmt.Println(fileInfo)
	for i := range data {
		// fmt.Println(k, v)
		if data[i].URL != "" {
			data[i].URL = fileInfo.Host + data[i].URL
		}
	}

	s.SuccessJSONDataPage(c, count, data)
}

// Create is
func (s StorageController) Create(c *gin.Context) {
	Storage := &model.Storage{}

	if err := c.BindJSON(Storage); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	err := s.Storage.Create(Storage)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, Storage)
}

// Update is
func (s StorageController) Update(c *gin.Context) {
	data := &model.Storage{}

	if err := c.BindJSON(data); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if data.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data2, err := s.Storage.Get(data.ID)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	data2.Name = data.Name

	err = data2.Update()
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONUpdate(c)
}

// Delete is
func (s StorageController) Delete(c *gin.Context) {
	category := &model.Storage{}

	if err := c.BindJSON(category); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if category.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data, err := s.Storage.Get(category.ID)
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
