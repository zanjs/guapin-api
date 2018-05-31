package controllers

import (
	"fmt"
	"mugg/guapin/app/middleware/jwtauth"
	"mugg/guapin/app/service"
	"mugg/guapin/model"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	// "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type (
	// UserController is
	UserController struct {
		BaseController
		User service.User
	}
)

// NewUser is
func NewUser() *UserController {
	return &UserController{}
}

// Home is
func (s UserController) Home(c *gin.Context) {
	users, _ := s.User.GetAll()

	s.SuccessJSONData(c, users)
}

// GetMe is
func (s UserController) GetMe(c *gin.Context) {
	useCla := c.MustGet("user").(*jwtauth.CustomClaims)

	fmt.Println(useCla.Name)

	u, err := s.User.GetByUsername(useCla.Name)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	fmt.Println(u)

	u.AvatarURL = "http://chuangyiren.cn/images/tmp/works_20150810_7031/2015081015413512303.jpg"

	s.SuccessJSONData(c, u)
}

// Create is
func (s UserController) Create(c *gin.Context) {
	userLogin := &model.UserLogin{}

	if err := c.BindJSON(userLogin); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if userLogin.Name == "" || userLogin.Password == "" {
		s.ErrorJSON(c, "name or password is null")
		return
	}

	useCla := c.MustGet("user").(*jwtauth.CustomClaims)

	fmt.Println(useCla.Name)

	user := &model.User{}
	user.Name = userLogin.Name
	user.Password = userLogin.Password

	err := s.User.Create(user)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	s.SuccessJSONData(c, user)
}

// Update is
func (s UserController) Update(c *gin.Context) {
	userLogin := &model.UserLogin{}

	if err := c.BindJSON(userLogin); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if userLogin.OldPassword == "" {
		s.ErrorJSON(c, "密码不能为空")
		return
	}

	if userLogin.OldPassword == userLogin.Password {
		s.SuccessJSONUpdate(c)
		return
	}

	useCla := c.MustGet("user").(*jwtauth.CustomClaims)

	fmt.Println(useCla.Name)
	fmt.Println(userLogin.Name)

	if useCla.Name != userLogin.Name {
		s.ErrorJSON(c, "非法操作")
		return
	}

	u, err := s.User.GetByUsername(userLogin.Name)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	fmt.Println(u)

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userLogin.OldPassword))

	if err != nil {
		s.ErrorJSON(c, "原密码错误")
		return
	}

	s.SuccessJSONUpdate(c)
}

// Login is
func (s UserController) Login(c *gin.Context) {
	user := &model.UserLogin{}

	if err := c.BindJSON(user); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if user.Name == "" || user.Password == "" {
		s.ErrorJSON(c, "name or password is null")
		return
	}

	u, err := s.User.GetByUsername(user.Name)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}
	fmt.Println(u)
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))

	if err != nil {
		s.ErrorJSON(c, "用户名或密码错误")
		return
	}

	j := &jwtauth.JWT{}
	j.SigningKey = []byte(jwtauth.SignKey)

	claims := jwtauth.CustomClaims{}

	claims.ID = u.ID
	claims.Name = u.Name
	claims.StandardClaims = jwt.StandardClaims{ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), //time.Now().Add(24 * time.Hour).Unix()
		Issuer: jwtauth.SignKey}

	token, err := j.CreateToken(claims)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	// _, err = j.ParseToken(token)

	// if err != nil {
	// 	if err == jwtauth.TokenExpired {
	// 		newToken, err := j.RefreshToken(token)
	// 		if err != nil {
	// 			s.ErrorJSON(c, err.Error())
	// 		} else {
	// 			s.SuccessJSONData(c, newToken)
	// 		}
	// 		return
	// 	}
	// 	s.ErrorJSON(c, err.Error())
	// 	return
	// }

	s.SuccessJSONData(c, token)
}
