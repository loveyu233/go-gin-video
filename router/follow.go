package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/api"
	"go-gin-video/middleware"
)

func FollowRouterInit(c *gin.Engine) {
	followApi := api.ApiGroupApp.FollowApi
	follow := c.Group("/follow")
	follow.Use(middleware.Auth)
	{
		follow.GET("is", followApi.IsFollow)
		follow.GET("add", followApi.AddFollow)
		follow.GET("del", followApi.DelFollow)
		follow.GET("followmeuser", followApi.GetFollowMeUser)
		follow.GET("mefollowcount", followApi.GetMeFollowCount)
		follow.GET("mefollowuser", followApi.GetMeFollowUser)
	}
	c.GET("/follow/followmecount", followApi.GetFollowMeCount)
}
