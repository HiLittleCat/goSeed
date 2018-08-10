package controller

import (
	"github.com/HiLittleCat/core"
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
	group := core.Routers.Group("user")
	group.POST("/create", UserC.Create)
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
	mobile := u.StrLength("手机号码", "mobile", 11)
	name := u.StrLenRange("名字", "name", 2, 6)
	logo := u.StrLenRange("头像url", "logo", 0, 50)
	userModel, err := userService.Create(mobile, name, logo)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Ok(userModel)
}
