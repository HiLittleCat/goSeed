package controller

import (
	"github.com/HiLittleCat/goSeed/model"

	"github.com/HiLittleCat/core"
)

// User controller
type User struct {
	core.Controller
}

/**
 * @api {post} /user/create 创建用户
 * @apiGroup User
 *
 * @apiParam {String} mobile 手机号码.
 * @apiParam {String} name 名字.
 * @apiParam {String} logo 头像.

 * @apiUse Res
 */
func (u *User) PostCreate() (*model.User, error) {
	mobile := u.ParamLength("mobile", 11)
	name := u.ParamGet("name")
	logo := u.ParamGet("logo")
	user := &model.User{Mobile: mobile, Name: name, Logo: logo}
	if err := user.Create(); err != nil {
		return nil, err
	}
	return user, nil
}

/**
 * @api {get} /user/ 获取用户信息
 * @apiGroup User
 *
 * @apiParam {String} _id 用户标识.

 * @apiUse Res
 */
func (u *User) Get() (*model.User, error) {
	_id := u.ParamLength("_id", 24)

	user := &model.User{ID: _id}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

/**
 * @api {get} /user/page  按页获取用户列表
 * @apiGroup User
 *
 * @apiParam {String} page 页码.

 * @apiUse Res
 */
func (u *User) GetPage() (*model.UserList, error) {
	page := u.ParamMin("page", 1)
	pageCount := u.ParamRange("pageCount", 10, 20)

	list := model.UserList{Page: page, PageCount: pageCount}
	err := list.GetPage()
	if err != nil {
		return nil, err
	}
	return &list, nil
}
