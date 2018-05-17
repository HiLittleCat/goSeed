package controller

import (
	"github.com/HiLittleCat/core"
	"github.com/HiLittleCat/goSeed/middleware"
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
	group.GET("/page/:number/:count", UserC.GetPage)
	group.GET("/get/:id", middleware.RequireSession, UserC.Get)
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
	_id := u.StrLength(ctx, "id", 24)

	userModel, err := userService.GetInfo(_id)
	if err != nil {
		ctx.Fail(err)
	} else {
		ctx.Ok(userModel)
	}
}

/**
 * @api {get} /user/page/:number/:count  按页获取用户列表
 * @apiGroup User
 * @apiParam {String} :number 页码.
 * @apiParam {String} :count 每页数据条数.
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

// login
// st := (&session.Store{}).Generate(_id, map[string]string{"uid": _id, "name": "xuyunfeng"})
// st.Flush()
// st.Cookie.Value = _id
// http.SetCookie(ctx.ResponseWriter, &st.Cookie)
