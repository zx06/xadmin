package service

import (
	"context"
	"reflect"
	"testing"
	"xadmin/app/model"
	"xadmin/app/serializer"
	"xadmin/app/serializer/request"
)

func TestUserLoginService_checkPassword(t *testing.T) {
	type args struct {
		user     *model.User
		password string
	}
	tests := []struct {
		name string
		u    *UserLoginService
		args args
		want bool
	}{
		{
			name: "test working",
			u:    &UserLoginService{},
			args: args{
				user: &model.User{
					PasswordDigest: "",
				},
				password: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserLoginService{}
			if got := u.checkPassword(tt.args.user, tt.args.password); got != tt.want {
				t.Errorf("UserLoginService.checkPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLoginService_Login(t *testing.T) {
	type args struct {
		ctx context.Context
		req request.UserLoginRequest
	}
	tests := []struct {
		name string
		u    *UserLoginService
		args args
		want serializer.Response
	}{
		{
			name: "",
			u:    &UserLoginService{},
			args: args{
				req: request.UserLoginRequest{
					Account:  "",
					Password: "",
				},
			},
			want: serializer.Response{
				Code:  40001,
				Data:  nil,
				Msg:   "账号或密码错误",
				Error: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserLoginService{}
			if got := u.Login(tt.args.ctx, tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserLoginService.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
