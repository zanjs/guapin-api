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
	c.JSON(200, gin.H{
		"va": "123",
	})
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
