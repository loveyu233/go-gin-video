package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/api"
	"go-gin-video/middleware"
)

func CommentRouter(c *gin.Engine) {
	authGroup := c.Group("")
	authGroup.Use(middleware.Auth)
	commentApi := api.ApiGroupApp.CommentApi
	comment := authGroup.Group("/comment")
	{
		comment.POST("/add", commentApi.InsertComment)
		comment.GET("/delete", commentApi.DeleteComment)
	}
	reply := authGroup.Group("/reply")
	{
		reply.POST("/add", commentApi.InsertReply)
		reply.GET("/delete", commentApi.DelReply)
	}

	{
		c.GET("/comment/get", commentApi.GetComment)
		c.GET("/comment/byid", commentApi.ByIdComment)
	}
}
