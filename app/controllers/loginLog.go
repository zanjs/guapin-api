package controllers

import (
	"mugg/guapin/app/middleware"
	"mugg/guapin/app/service"
	"mugg/guapin/model"

	"github.com/gin-gonic/gin"
)

type (
	// LoginLogController is
	LoginLogController struct {
		BaseController
		LoginLog service.LoginLog
	}
)

// NewLoginLog is
func NewLoginLog() *LoginLogController {
	return &LoginLogController{}
}

// Home is
func (s LoginLogController) Home(c *gin.Context) {
	qPage := middleware.GetPage(c)

	searchName := c.DefaultQuery("title", "")

	data, err := s.LoginLog.GetAllQuerySearch(qPage, searchName)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	count, err := s.LoginLog.GetAllQuerySearchTotal(searchName)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	s.SuccessJSONDataPage(c, count, data)
}

// Create is
func (s LoginLogController) Create(c *gin.Context) {
	LoginLog := &model.LoginLog{}

	if err := c.BindJSON(LoginLog); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	err := s.LoginLog.Create(LoginLog)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, LoginLog)
}

// Update is
func (s LoginLogController) Update(c *gin.Context) {
	data := &model.LoginLog{}

	if err := c.BindJSON(data); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if data.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data2, err := s.LoginLog.Get(data.ID)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	data2.IP = data.IP
	data2.Name = data.Name
	data2.UserAgent = data.UserAgent

	err = data2.Update()
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONUpdate(c)
}

// Delete is
func (s LoginLogController) Delete(c *gin.Context) {
	data := &model.LoginLog{}

	if err := c.BindJSON(data); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if data.ID == 0 {
		s.ErrorJSON(c, "")
		return
	}

	data2, err := s.LoginLog.Get(data.ID)
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
