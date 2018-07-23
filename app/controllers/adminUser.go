package controllers

import (
	"fmt"
	"mugg/guapin/app/conf"
	"mugg/guapin/app/middleware"
	"mugg/guapin/app/middleware/jwtauth"
	"mugg/guapin/app/service"
	"mugg/guapin/model"
	"mugg/guapin/utils"

	// "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type (
	// AdminUserController is
	AdminUserController struct {
		BaseController
		AdminUser     service.AdminUser
		AdminLoginLog service.AdminLoginLog
	}
)

// NewAdminUser is
func NewAdminUser() *AdminUserController {
	return &AdminUserController{}
}

// Home is
func (s AdminUserController) Home(c *gin.Context) {
	qPage := middleware.GetPage(c)

	searchName := c.DefaultQuery("title", "")

	data, err := s.AdminUser.GetAllQuerySearch(qPage, searchName)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	count, err := s.AdminUser.GetAllQuerySearchTotal(searchName)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	for i := range data {
		if data[i].Avatar == "" {
			data[i].Avatar = conf.Config.AdminUser.Avatar
		}
	}

	s.SuccessJSONDataPage(c, count, data)
}

// GetMe is
func (s AdminUserController) GetMe(c *gin.Context) {
	fmt.Println("\n")
	fmt.Println("进来啦")
	fmt.Println("\n")

	useCla := c.MustGet("user").(*jwtauth.CustomClaims)

	fmt.Println(useCla.Name)

	u, err := s.AdminUser.GetByAdminUsername(useCla.Name)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	fmt.Println(u)

	u.Avatar = conf.Config.AdminUser.Avatar
	s.SuccessJSONData(c, u)
}

// Create is
func (s AdminUserController) Create(c *gin.Context) {
	userLogin := &model.AdminUserLogin{}

	if err := c.BindJSON(userLogin); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if userLogin.Name == "" || userLogin.Password == "" {
		s.ErrorJSON(c, "name or password is null")
		return
	}

	// useCla := c.MustGet("user").(*jwtauth.CustomClaims)

	// fmt.Println(useCla.Name)

	user := &model.AdminUser{}
	user.Name = userLogin.Name
	user.Password = userLogin.Password
	user.Avatar = conf.Config.AdminUser.Avatar

	err := s.AdminUser.Create(user)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	s.SuccessJSONData(c, user)
}

// Update is
func (s AdminUserController) Update(c *gin.Context) {
	data := &model.AdminUserUpdate{}

	if err := c.BindJSON(data); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	// if userLogin.OldPassword == userLogin.Password {
	// 	s.SuccessJSONUpdate(c)
	// 	return
	// }

	useCla := c.MustGet("user").(*jwtauth.CustomClaims)

	fmt.Println(useCla.Name)
	fmt.Println(data.Name)

	// if useCla.Name != userLogin.Name {
	// 	s.ErrorJSON(c, "非法操作")
	// 	return
	// }

	u, err := s.AdminUser.GetByAdminUsername(data.Name)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	fmt.Println(u)

	// err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userLogin.OldPassword))

	// if err != nil {
	// 	s.ErrorJSON(c, "原密码错误")
	// 	return
	// }
	if data.Password != "" {
		u.Password = utils.HashPassword(data.Password)
		u.Update()
	}

	s.SuccessJSONUpdate(c)
}
