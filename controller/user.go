package controller

import (
	"github.com/HiLittleCat/goSeed/model"

	"github.com/HiLittleCat/core"
)

// User controller
type User struct {
	core.Controller
}

// PostCreate create user
func (u *User) PostCreate() (string, error) {
	mobile := u.ParamLength("mobile", 11)
	name := u.ParamGet("name")
	logo := u.ParamGet("logo")
	user := model.User{Mobile: mobile, Name: name, Logo: logo}
	if err := user.Create(); err != nil {
		return "", err
	}
	return "create user ok", nil
}

// Get User
func (u *User) Get() (*model.User, error) {
	_id := u.ParamLength("_id", 24)
	user := model.User{ID: _id}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetPage User/page
func (u *User) GetPage() (*model.UserList, error) {
	page := u.ParamMin("page", 1)
	pageCount := u.ParamMin("pageCount", 10)
	list := model.UserList{Page: page, PageCount: pageCount}
	err := list.GetPage()
	if err != nil {
		return nil, err
	}
	return &list, nil
}
