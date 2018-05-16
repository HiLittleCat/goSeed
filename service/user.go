package service

import (
	"github.com/HiLittleCat/core"
	"github.com/HiLittleCat/goSeed/errors"
	"github.com/HiLittleCat/goSeed/model"
)

// User service
type User struct {
	core.Service
}

//Create  获取用户信息
func (u *User) Create(mobile string, name string, logo string) (*model.User, error) {
	userModel := &model.User{Mobile: mobile, Name: name, Logo: logo}

	//检查用户是否存在
	count, err := userModel.GetCountByID()
	if count > 0 {
		return nil, errors.ErrUserExist
	} else if err != nil {
		return nil, err
	}

	//创建用户
	err = userModel.Create()

	return userModel, err
}

//GetInfo  获取用户信息
func (u *User) GetInfo(_id string) (*model.User, error) {
	userModel := &model.User{ID: _id}
	err := userModel.GetByID()
	return userModel, err
}

//GetPage  按页获取用户列表
func (u *User) GetPage(page int, pageCount int) (*model.UserList, error) {
	list := model.UserList{Page: page, PageCount: pageCount}
	err := list.GetPage()
	return &list, err
}
