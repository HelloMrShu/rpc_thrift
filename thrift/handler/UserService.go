package handler

import (
	"context"
	"user/thrift/user"
)

type UserService struct {
}

// 实现IDL里定义的接口

func (us UserService) GetName(ctx context.Context, id int32) (name string, err error) {

	u := user.UserInfo{}
	u.Name = "gao qi"

	return u.Name, nil
}

func (us UserService) GetUser(ctx context.Context, id int32) (ui *user.UserInfo, err error) {

	u := &user.UserInfo{}
	u.Name = "gao qi"

	return u, nil
}
