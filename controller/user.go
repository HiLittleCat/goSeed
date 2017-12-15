package controller

import (
	"strconv"

	"github.com/HiLittleCat/goSeed/model"

	"net/http"

	"github.com/HiLittleCat/core"
)

type User struct {
	core.Controller
}

/**
 */
func (u *User) PostCreate() {
	mobile := u.Ctx.Request.FormValue("mobile")
	if mobile == "" {
		u.Ctx.Fail(http.StatusBadRequest, "User.Create post param 'mobile' is required.")
		return
	}
	name := u.Ctx.Request.FormValue("name")
	logo := u.Ctx.Request.FormValue("logo")
	user := model.User{Mobile: mobile, Name: name, Logo: logo}
	if err := user.Create(); err != nil {
		u.Ctx.Fail(http.StatusInternalServerError, "User.Create error.", err)
		return
	}
}

/**
* User
 */

func (u *User) Get() {
	_id := u.Ctx.Request.FormValue("_id")

	if _id == "" {
		u.Ctx.Fail(http.StatusBadRequest, "User.Get query param '_id' is required.")
		return
	}
	user := model.User{Id: _id}
	if err := user.Get(); err != nil {
		u.Ctx.Fail(http.StatusInternalServerError, "User.Get err.", err)
		return
	}
	u.Ctx.Success(http.StatusOK, user)
}

/**
* User/page
 */
func (u *User) GetPage() {
	page, err := strconv.Atoi(u.Ctx.Request.FormValue("page"))
	if err != nil {
		u.Ctx.Fail(http.StatusBadRequest, "User.GetPage query param 'page' must be a number.")
		return
	}
	pageCount, err := strconv.Atoi(u.Ctx.Request.FormValue("pageCount"))
	if err != nil {
		u.Ctx.Fail(http.StatusBadRequest, "User.GetPage query param 'pageCount' must be a number.")
		return
	}
	list := model.UserList{Page: page, PageCount: pageCount}
	err = list.GetPage()
	if err != nil {
		u.Ctx.Fail(http.StatusInternalServerError, "User.GetPage fail to get list.", err)
		return
	}
	u.Ctx.Success(http.StatusOK, list)
}
