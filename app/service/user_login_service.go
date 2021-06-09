package service

import (
	"context"
	"time"
	"xadmin/app/config"
	"xadmin/app/model"
	"xadmin/app/serializer"
	"xadmin/app/serializer/request"

	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	Account string `json:"account"`
	Admin   bool   `json:"admin"`
	jwt.StandardClaims
}

type UserLoginService struct {
}

// checkPassword 检查用的密码是否匹配
func (u *UserLoginService) checkPassword(user *model.User, password string) bool {
	// todo 检查逻辑
	return true
}

// Login 登录服务
func (u *UserLoginService) Login(ctx context.Context, req request.UserLoginRequest) serializer.Response {
	var user model.User
	if err := config.DB().WithContext(ctx).Where("account = ?", req.Account).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}
	if !u.checkPassword(&user, req.Password) {
		return serializer.ParamErr("账号或密码错误", nil)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtClaims{
		Account: user.Account,
		Admin:   user.Admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	})
	t, err := token.SignedString(config.Config().SigningKey)
	if err != nil {
		return serializer.Err(serializer.CodeEncryptError, "token 生成失败", nil)
	}
	return serializer.Response{
		Data:  t,
	}
}
