package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/api"
	"go-gin-video/middleware"
)

func LikeRouterInit(c *gin.Engine) {
	likeApi := api.ApiGroupApp.LikeApi
	like := c.Group("/like")
	like.Use(middleware.Auth)
	{
		like.GET("/add", likeApi.AddLike)
		like.GET("/del", likeApi.DelLike)
		like.GET("/is", likeApi.IsLike)
	}
	c.GET("/like/count", likeApi.LikeCount)
}
