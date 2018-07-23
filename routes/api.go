package routes

import (
	"mugg/guapin/app/controllers"
	"mugg/guapin/routes/api"

	"github.com/gin-gonic/gin"
)

// API is router api
func API(g *gin.Engine) {
	catController := controllers.NewCat()

	apiG := g.Group("/api")

	v1 := apiG.Group("/v1")
	{
		v1.GET("/", catController.Home)

		// v1.GET("/user/:name/:action", func(c *gin.Context) {
		// 	name := c.Param("name")
		// 	action := c.Param("action")
		// 	message := name + " is " + action
		// 	c.String(http.StatusOK, message)
		// })
	}

	api.AdminUser(v1)
	api.Table(v1)
	api.Category(v1)
	api.Page(v1)
	api.Upload(v1)
	api.Article(v1)
	api.Storage(v1)
	api.Setting(v1)
	api.Logs(v1)
	api.LoginLog(v1)
	api.Excel(v1)

	// banner
	api.Banner(v1)
	api.BannerCategory(v1)

	// goods
	api.Goods(v1)
	api.GoodsCategory(v1)

}
