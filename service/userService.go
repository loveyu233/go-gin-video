package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-gin-video/domain/dto"
	"go-gin-video/service/mysqlService"
	"go-gin-video/service/redisService"
	"go-gin-video/utils/email"
	"go-gin-video/utils/jwt"
	"go-gin-video/utils/md5Encodeing"
	"go-gin-video/utils/strconv"
	"go-gin-video/utils/upload"
)

type UserServices struct {
}

func (UserServices) UserLogin(userDto *dto.UserDto) (error, string) {
	userDto.Password = md5Encodeing.PasswordMd5Encryption(userDto.Password)
	userModel, err := mysqlService.User.Where(mysqlService.User.Username.Eq(userDto.Username), mysqlService.User.Password.Eq(userDto.Password)).First()
	if err != nil {
		return err, ""
	}
	token, err := jwt.CreateToken(userModel.UserID, userModel.Role)
	if err != nil {
		return err, ""
	}
	return nil, token
}

func (UserServices) UserRegister(userDto *dto.UserDto) error {
	if userDto.Username == "" || userDto.Password == "" || userDto.Email == "" || userDto.EmailCode == "" {
		return errors.New("缺少参数")
	}
	if !email.VerifyEmailFormat(userDto.Email) {
		return errors.New("邮箱格式错误")
	}
	if !redisService.VerificationEmailCode(userDto.Email, userDto.EmailCode) {
		return errors.New("邮箱验证码错误")
	}
	userModel := userDto.ToUserModel()
	err := mysqlService.User.Create(userModel)
	if err != nil {
		return errors.New("用户创建失败")
	}
	return nil
}

func (UserServices) UserResetPassword(userDto *dto.UserDto) error {
	if !email.VerifyEmailFormat(userDto.Email) {
		return errors.New("邮箱格式错误")
	}
	if !redisService.VerificationEmailCode(userDto.Email, userDto.EmailCode) {
		return errors.New("邮箱验证码错误")
	}
	userModel := userDto.ToUserModel()
	_, err := mysqlService.User.Where(mysqlService.User.Email.Eq(userModel.Email)).Update(mysqlService.User.Password, userModel.Password)
	if err != nil {
		return errors.New("用户信息保存失败")
	}
	return nil
}

func (UserServices) UpdateUser(userDto *dto.UserDto, c *gin.Context) error {
	if !email.VerifyEmailFormat(userDto.Email) {
		return errors.New("邮箱格式错误")
	}
	if !redisService.VerificationEmailCode(userDto.Email, userDto.EmailCode) {
		return errors.New("验证码错误")
	}

	if !upload.IsImageOverflow(&userDto.UserIcon, &userDto.UserBg) {
		return errors.New("上传文件超出大小限制")
	}

	is, icon, bg := upload.SaveUserIconAndBg(c, userDto)
	if !is {
		return errors.New("用户上传文件保存失败")
	}
	userModel := userDto.ToUserModelUpdate(icon, bg)
	queryUser := mysqlService.User
	userid := strconv.StringToInt(c.GetString("userid"))
	_, err := queryUser.Where(queryUser.UserID.Eq(int32(userid))).UpdateColumns(userModel)
	if err != nil {
		return err
	}
	return nil
}
