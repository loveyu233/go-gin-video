package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/api"
)

func CaptchaRouter(c *gin.Engine) {
	captcha := c.Group("captcha")
	captchaApi := api.ApiGroupApp
	{
		captcha.GET("/get", captchaApi.GetCaptchaCode)
		captcha.GET("/ver", captchaApi.VerificationCaptchaCode)
	}
}
