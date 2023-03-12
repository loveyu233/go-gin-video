package dto

import (
	"go-gin-video/domain/model"
	"go-gin-video/domain/vo"
	"go-gin-video/utils/md5Encodeing"
	"mime/multipart"
)

type UserDto struct {
	Username  string               `json:"username"`  // 用户名
	Password  string               `json:"password"`  //用户密码
	Email     string               `json:"email"`     //用户邮箱
	EmailCode string               `json:"emailCode"` //邮箱验证码
	UserIcon  multipart.FileHeader `json:"user_icon"` // 用户头像
	Sing      string               `json:"sing"`      // 个性签名
	UserBg    multipart.FileHeader `json:"user_bg"`   // 用户主页背景
}

func (u UserDto) ToUserModelUpdate(userIconPath, userBgPath string) *model.User {
	return &model.User{
		Username: u.Username,
		UserIcon: userIconPath,
		Sing:     u.Sing,
		UserBg:   userBgPath,
	}
}

func (u *UserDto) ToUserModel() *model.User {
	return &model.User{
		Username: u.Username,
		Password: md5Encodeing.PasswordMd5Encryption(u.Password),
		Email:    u.Email,
	}
}

func UserModelToUserVo(userModel *model.User) *vo.UserVo {
	return &vo.UserVo{
		Username: userModel.Username,
		UserIcon: userModel.UserIcon,
		Sing:     userModel.Sing,
		UserBg:   userModel.UserBg,
	}
}
