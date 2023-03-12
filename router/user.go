package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/api"
)

func UserRouter(c *gin.Engine) {
	user := c.Group("/user")
	userApi := api.ApiGroupApp.UserApi
	{
		user.POST("/login", userApi.UserLogin)
		user.GET("/authcode", userApi.UserSendAuthEmailCoe)
		user.POST("/register", userApi.UserRegister)
		user.POST("/resetpassword", userApi.ResetPassword)
		user.GET("/byid", userApi.ById)
		user.POST("/update", userApi.UpdateUser)
	}
}
