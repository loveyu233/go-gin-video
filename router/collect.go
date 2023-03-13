package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/api"
	"go-gin-video/middleware"
)

func CollectRouterInit(c *gin.Engine) {
	collectApi := api.ApiGroupApp.CollectApi
	collect := c.Group("/collect")
	collect.Use(middleware.Auth)
	{
		collect.GET("/is", collectApi.IsCollect)
		collect.GET("/add", collectApi.AddCollect)
		collect.GET("/del", collectApi.DelCollect)
		collect.GET("/usercollect", collectApi.GetUserCollect)
	}
	{
		c.GET("/collect/count", collectApi.CollectCount)
	}
}
