package mysql

import "errors"

var (
	ErrorUserExist           = errors.New("用户已经存在")
	ErrorUserNotExist        = errors.New("用户不存在")
	ErrorInvalidPassword     = errors.New("密码错误")
	ErrorInvalidID           = errors.New("无效的ID")
	ErrorUserLeave           = errors.New("用户离职")
	ErrorUserPasswordExpired = errors.New("密码的有效期过了")
	UserPasswordErrorGt3     = errors.New("24小时密码错误超过3次")
)
