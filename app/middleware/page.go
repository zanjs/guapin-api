package middleware

import (
	"mugg/guapin/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPage 获取分页数
func GetPage(c *gin.Context) model.QueryParamsPage {

	pageNoStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("limit", "20")
	var pageNo int
	var pageSize int
	var err error
	if pageNo, err = strconv.Atoi(pageNoStr); err != nil {
		pageNo = 1
	}

	if pageSize, err = strconv.Atoi(pageSizeStr); err != nil {
		pageSize = 20
	}

	if pageNo < 1 {
		pageNo = 1
	}

	if pageSize < 1 {
		pageSize = 20
	}

	queryPage := model.QueryParamsPage{}

	queryPage.Page = pageNo

	queryPage.Limit = pageSize

	queryPage.Offset = (queryPage.Page - 1) * pageSize

	return queryPage
}
