package controllers

import (
	"github.com/gin-gonic/gin"
)

type errorCode struct {
	SUCCESS      int
	ERROR        int
	AuthErr      int
	RoleErr      int
	NotFound     int
	LoginError   int
	LoginTimeout int
	InActive     int
}

// ErrorCode 错误码
var ErrorCode = errorCode{
	SUCCESS:      0,
	ERROR:        1,
	AuthErr:      40029, // 认证失败，请重新登陆
	RoleErr:      1406,  // 权限不够
	NotFound:     404,
	LoginError:   1000, //用户名或密码错误
	LoginTimeout: 1001, //登录超时
	InActive:     1002, //未激活账号
}

// JWTError is
func (b BaseController) JWTError(c *gin.Context, str string) {
	b.ErrorJSON(c, "认证失败，请重新登陆")
}

// Claims is
type Claims map[string]interface{}

// JWTHandler is
// func (b BaseController) JWTHandler(c *gin.Context) {
// 	response.JSON(ctx, user)
// }
