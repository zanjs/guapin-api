package controllers

import (
	"fmt"
	"mugg/guapin/app/middleware/jwtauth"
	"mugg/guapin/model"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	// "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login is
func (s AdminUserController) Login(c *gin.Context) {
	user := &model.AdminUserLogin{}

	if err := c.BindJSON(user); err != nil {
		s.ErrorJSON(c, err.Error())
		return
	}

	if user.Name == "" || user.Password == "" {
		s.ErrorJSON(c, "name or password is null")
		return
	}

	u, err := s.AdminUser.GetByAdminUsername(user.Name)
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
	userLoginLog := &model.AdminLoginLog{}

	userLoginLog.IP = ip
	userLoginLog.Name = u.Name
	userLoginLog.AdminUserID = u.ID
	userLoginLog.AdminUserAgent = userAgent

	// limiter := time.Tick(time.Millisecond * 2000)

	go func() {
		// <-limiter
		err = s.AdminLoginLog.Create(userLoginLog)
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
