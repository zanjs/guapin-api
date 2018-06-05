package controllers

import (
	"mugg/guapin/app/service"
	"mugg/guapin/model"

	"github.com/gin-gonic/gin"
)

type (
	// SettingController is
	SettingController struct {
		BaseController
		Setting service.Setting
	}
)

// NewSetting is
func NewSetting() *SettingController {
	return &SettingController{}
}

// Home is
func (s SettingController) Home(c *gin.Context) {
	// qPage := middleware.GetPage(c)

	// data, err := s.Setting.GetAllQuery(qPage)
	// if err != nil {
	// 	s.ErrorJSON(c, err.Error())
	// 	return
	// }
	// count, err := s.Setting.GetAllQueryTotal()
	// if err != nil {
	// 	s.ErrorJSON(c, err.Error())
	// 	return
	// }

	data, err := s.Setting.GetAll()
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	// s.SuccessJSONDataPage(c, count, data)
	s.SuccessJSONData(c, data)
}

// Create is
func (s SettingController) Create(c *gin.Context) {
	Setting := &model.Setting{}

	if err := c.BindJSON(Setting); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	err := s.Setting.Create(Setting)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, Setting)
}

// Update is
func (s SettingController) Update(c *gin.Context) {
	data := &model.Setting{}

	if err := c.BindJSON(data); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if data.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data2, err := s.Setting.Get(data.ID)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	data2.Key = data.Key
	data2.Value = data.Value
	data2.Name = data.Name
	data2.Type = data.Type

	err = data2.Update()
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONUpdate(c)
}

// Delete is
func (s SettingController) Delete(c *gin.Context) {
	data := &model.Setting{}

	if err := c.BindJSON(data); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if data.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data2, err := s.Setting.Get(data.ID)
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
