package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/domain/dto"
	"go-gin-video/domain/resp"
	"go-gin-video/service"
	"go-gin-video/service/mysqlService"
	"go-gin-video/service/redisService"
	"go-gin-video/utils/email"
	"go-gin-video/utils/strconv"
	"go.uber.org/zap"
)

type UserApi struct {
}

var userService = service.ServiceApp

// @Summary 用户登录
// @Tags user
// @Router /user/login [post]
func (UserApi) UserLogin(c *gin.Context) {
	var userDto *dto.UserDto
	err := c.Bind(&userDto)
	if err != nil {
		resp.ErrorResponse(c, "参数有误")
		return
	}
	err, token := userService.UserLogin(userDto)
	if err != nil {
		resp.ErrorResponse(c, "账号或密码错误")
		return
	}
	resp.OkResponse(c, token)
}

// @Summary 用户邮箱验证码发送
// @Tags user
// @Router /user/authcode [get]
func (UserApi) UserSendAuthEmailCoe(c *gin.Context) {
	emailStr := c.Query("email")
	if !email.VerifyEmailFormat(emailStr) {
		resp.ErrorResponse(c, "邮箱格式不正确请重新尝试")
		return
	}
	err := redisService.SaveEmailCode(emailStr)
	if err != nil {
		zap.L().Error(err.Error())
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, "验证码发送成功")
}

// @Summary 用户注册
// @Tags user
// @Router /user/register [post]
func (UserApi) UserRegister(c *gin.Context) {
	var user *dto.UserDto
	err := c.Bind(&user)
	if err != nil {
		zap.L().Error(err.Error())
		resp.ErrorResponse(c, "参数绑定失败")
		return
	}
	err = userService.UserRegister(user)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, "用户创建成功")
}

// @Summary 重置密码
// @Tags user
// @Router /user/resetpassword [post]
func (UserApi) ResetPassword(c *gin.Context) {
	var userDto *dto.UserDto
	err := c.Bind(userDto)
	if err != nil {
		zap.L().Error(err.Error())
		resp.ErrorResponse(c, "用户参数有误")
		return
	}
	err = userService.UserResetPassword(userDto)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, "密码修改成功")
}

// @Summary 通过userid获取用户信息
// @Tags user
// @Router /user/byid [get]
func (UserApi) ById(c *gin.Context) {
	queryUser := mysqlService.User
	userid := c.DefaultQuery("userid", "1")
	userModel, err := queryUser.Select(queryUser.UserID, queryUser.Username, queryUser.UserIcon).Where(queryUser.UserID.Eq(int32(strconv.StringToUint(userid)))).First()
	if err != nil {
		resp.ErrorResponse(c, "用户不存在")
		return
	}
	userVo := dto.UserModelToUserVo(userModel)
	resp.OkResponse(c, userVo)
}

// @Summary 修改用户信息
// @Tags user
// @Router /user/update [post]
func (UserApi) UpdateUser(c *gin.Context) {
	userDto := &dto.UserDto{}
	err := c.ShouldBind(userDto)
	if err != nil {
		zap.L().Error("用户信息绑定失败:" + err.Error())
		resp.ErrorResponse(c, "用户信息绑定失败")
		return
	}
	err = userService.UpdateUser(userDto, c)
	if err != nil {
		resp.ErrorResponse(c, err.Error())
		return
	}
	resp.OkResponse(c, "用户信息修改成功")
}
