package handler

import (
	"context"
	"errors"
	"user/thrift/user"
)

type UserService struct {
}

// 实现IDL里定义的接口

func (us UserService) GetName(ctx context.Context, id int32) (name string, err error) {
	if id < 0 {
		err = errors.New("id不能小于0")
		return
	}

	u := user.UserInfo{}
	u.Name = "gao qi"
	name = "gao qi"

	return
}

func (us UserService) GetUser(ctx context.Context, id int32) (ui *user.UserInfo, err error) {

	u := &user.UserInfo{}
	u.Name = "gao qi"

	return u, nil
}
