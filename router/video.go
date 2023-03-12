package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/api"
	"go-gin-video/middleware"
)

func VideoRouter(c *gin.Engine) {
	video := c.Group("/video")
	videoApi := api.ApiGroupApp.VideoApi
	{
		video.GET("/get", videoApi.ByIdVideo)
		video.GET("/list", videoApi.GetVideoList)
		video.GET("/search", videoApi.SearchVideoList)
		video.GET("/user", videoApi.UserIdFindVideo)
	}
	{
		video.GET("/status", middleware.Auth, videoApi.SelectNotCheck)
	}
	root := c.Group("/root")
	root.Use(middleware.Auth)
	{
		root.GET("/video/status", videoApi.SelectNotCheckAll)
		root.GET("/video/check", videoApi.UpdateVideoCheck)
	}
}
