package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/api"
	"go-gin-video/middleware"
)

func UploadRouter(c *gin.Engine) {
	upload := c.Group("/upload")
	upload.Use(middleware.Auth)
	uploadApi := api.ApiGroupApp.UploadApi
	{
		upload.POST("/vide", uploadApi.UploadVideo)
	}
}
