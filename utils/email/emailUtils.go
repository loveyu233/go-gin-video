package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"regexp"
)

// 发送邮箱验证码
func EmailSendMsg(email, code string) bool {
	sender := viper.GetString("email.username")   //发送者qq邮箱
	authCode := viper.GetString("email.password") //qq邮箱授权码
	mailTitle := "验证码"                            //邮件标题
	mailBody := code                              //邮件内容,可以是html

	//接收者邮箱列表
	mailTo := []string{
		email,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", sender)       //发送者腾讯邮箱账号
	m.SetHeader("To", mailTo...)      //接收者邮箱列表
	m.SetHeader("Subject", mailTitle) //邮件标题
	m.SetBody("text/html", mailBody)  //邮件内容,可以是html

	// //添加附件
	// zipPath := "./user/temp.zip"
	// m.Attach(zipPath)

	//发送邮件服务器、端口、发送者qq邮箱、qq邮箱授权码
	//服务器地址和端口是腾讯的
	d := gomail.NewDialer("smtp.qq.com", 587, sender, authCode)
	err := d.DialAndSend(m)
	if err != nil {
		return false
	}
	return true
}

// 验证邮箱格式是否正确
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
