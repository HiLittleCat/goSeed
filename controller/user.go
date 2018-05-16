package controller

import (
	"github.com/HiLittleCat/core"
	"github.com/HiLittleCat/goSeed/errors"
	"github.com/HiLittleCat/goSeed/service"
)

var (
	userService *service.User
)

// User controller
type User struct {
	core.Controller
}

// Register register routers of this controller
func (u *User) Register() {
	UserC := &User{}
	rUser := core.Routers.Group("user")
	rUser.POST("/create", UserC.Create)
	rUser.GET("/page/:number/:count", UserC.GetPage)
	rUser.GET("/get/:id", UserC.Get)
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
func (u *User) Create(ctx *core.Context) {
	mobile := u.StrLength(ctx, "mobile", 11)
	name := u.StrGet(ctx, "name")
	logo := u.StrGet(ctx, "logo")
	userModel, err := userService.Create(mobile, name, logo)
	if err != nil {
		ctx.Fail(err)
	} else {
		ctx.Ok(userModel)
	}
}

/**
 * @api {get} /user/:id 获取用户信息
 * @apiGroup User
 *
 * @apiParam {String} :id 用户标识.

 * @apiUse Res
 */
func (u *User) Get(ctx *core.Context) {
	session := ctx.Data["session"]
	if session == nil {
		ctx.Fail(errors.ErrUserExist)
		return
	}
	_id := u.StrLength(ctx, "id", 1)
	userModel, err := userService.GetInfo(_id)
	if err != nil {
		ctx.Fail(err)
	} else {
		ctx.Ok(userModel)
	}
}

/**
 * @api {get} /user/page/:page  按页获取用户列表
 * @apiGroup User
 *
 * @apiParam {String} page 页码.

 * @apiUse Res
 */
func (u *User) GetPage(ctx *core.Context) {
	page := u.IntRange(ctx, "number", 1, 100)
	pageCount := u.IntRange(ctx, "count", 10, 20)
	list, err := userService.GetPage(page, pageCount)
	if err != nil {
		ctx.Fail(err)
	} else {
		ctx.Ok(list)
	}
}
