package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/domain/resp"
	"go-gin-video/service/redisService"
	"go-gin-video/utils/captcha"
	"go-gin-video/utils/email"
	"go-gin-video/utils/strconv"
)

type CaptchaApi struct {
}

// @Summary 获取滑块验证码
// @Tags captcha
// @Router /captcha/get [get]
func (CommentApi) GetCaptchaCode(c *gin.Context) {
	emailStr := c.Query("email")
	if !email.VerifyEmailFormat(emailStr) {
		resp.ErrorResponse(c, "邮箱格式错误")
		return
	}
	captchaModel := captcha.CreateCode()
	modelToVo := captcha.CaptchaModelToVo(captchaModel)
	redisService.SaveCaptchaCode(emailStr, captchaModel.X)
	resp.OkResponse(c, modelToVo)
}

// @Summary 验证滑块验证码
// @Tags captcha
// @Router /captcha/ver [get]
func (CommentApi) VerificationCaptchaCode(c *gin.Context) {
	emailStr := c.Query("email")
	if !email.VerifyEmailFormat(emailStr) {
		resp.ErrorResponse(c, "邮箱格式错误")
		return
	}
	xValue := c.Query("x")
	if !redisService.VerificationCaptchaCode(emailStr, strconv.StringToInt(xValue)) {
		resp.ErrorResponse(c, "验证码错误")
		return
	}
	resp.OkResponse(c, "验证通过")
}
