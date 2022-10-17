package mail

import (
	"gopkg.in/gomail.v2"
)

// https://blog.csdn.net/qq_41971768/article/details/109726349

const (
	// 邮件服务器地址
	MailHost = "smtp.163.com"
	// 端口
	MailPort = 25
	// 发送邮件用户账号[自己的]
	MailUser = "ljs_hsm@163.com"
	// 授权密码
	MailPwd = "RCHNMWKGWDSZLEXB"
)

/*
   title 使用gomail发送邮件
   @param []string mailAddress 收件人邮箱
   @param string subject 邮件主题
   @param string body 邮件内容
   @return error
*/
func SendGoMail(mailAddress []string, subject string, body string) error {
	m := gomail.NewMessage()
	// 这种方式可以添加别名，即 nickname， 也可以直接用<code>m.SetHeader("From", MAIL_USER)</code>
	nickname := "gomail"
	m.SetHeader("From", nickname+"<"+MailUser+">")
	// 发送给多个用户
	m.SetHeader("To", mailAddress...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(MailHost, MailPort, MailUser, MailPwd)
	// 发送邮件
	err := d.DialAndSend(m)
	return err
}
