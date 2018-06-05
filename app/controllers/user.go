package controllers

import (
	"fmt"
	"mugg/guapin/app/conf"
	"mugg/guapin/app/middleware/jwtauth"
	"mugg/guapin/app/service"
	"mugg/guapin/model"
	"mugg/guapin/utils"
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
		User     service.User
		LoginLog service.UserLoginLog
	}
)

// NewUser is
func NewUser() *UserController {
	return &UserController{}
}

// Home is
func (s UserController) Home(c *gin.Context) {
	data, _ := s.User.GetAll()

	for i := range data {
		if data[i].Avatar == "" {
			data[i].Avatar = conf.Config.User.Avatar
		}
	}

	s.SuccessJSONData(c, data)
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

	u.Avatar = conf.Config.User.Avatar
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

	// useCla := c.MustGet("user").(*jwtauth.CustomClaims)

	// fmt.Println(useCla.Name)

	user := &model.User{}
	user.Name = userLogin.Name
	user.Password = userLogin.Password
	user.Avatar = conf.Config.User.Avatar

	err := s.User.Create(user)
	if err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	s.SuccessJSONData(c, user)
}

// Update is
func (s UserController) Update(c *gin.Context) {
	data := &model.UserUpdate{}

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

	u, err := s.User.GetByUsername(data.Name)
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
	ip := c.ClientIP()
	userAgent := c.GetHeader("user-agent")
	userLoginLog := &model.UserLoginLog{}

	userLoginLog.IP = ip
	userLoginLog.Name = u.Name
	userLoginLog.UserID = u.ID
	userLoginLog.UserAgent = userAgent

	limiter := time.Tick(time.Millisecond * 2000)

	go func() {
		<-limiter
		err = s.LoginLog.Create(userLoginLog)
		if err != nil {
			fmt.Println(err.Error())
		}

		// sum := 0
		// for {
		// 	sum++
		// 	if sum > 100 {
		// 		break
		// 	}
		// 	<-limiter //执行到这里，需要隔 200毫秒才继续往下执行，time.Tick(timer)上面已定义
		// 	fmt.Println("request", sum, time.Now())
		// }
	}()

	s.SuccessJSONData(c, token)
}
