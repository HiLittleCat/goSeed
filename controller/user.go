package controller

import (
	"strconv"

	"github.com/HiLittleCat/goSeed/model"

	"net/http"

	"github.com/HiLittleCat/core"

	log "github.com/sirupsen/logrus"
)

type User struct {
	core.Controller
}

/**
* User
 */

func (userController *User) Get() *model.User {
	_id := userController.Ctx.QueryParam("_id")
	if _id == "" {
		log.Warnln("User.Get query param _id must is required.")
		userController.Ctx.ResStatus(http.StatusBadRequest)
		return nil
	}
	user := model.User{Id: _id}
	if err := user.Get(); err != nil {
		log.WithFields(log.Fields{"err": err}).Warnln("User.Get error")
		userController.Ctx.ResStatus(http.StatusInternalServerError)
		return nil
	}
	return &user
}

/**
* User/page
 */
func (userController *User) GetPage() []model.User {
	user := model.User{}
	page, err := strconv.Atoi(userController.Ctx.QueryParam("page"))
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warnln("User.GetPage query param page must be a number.")
		userController.Ctx.ResStatus(http.StatusBadRequest)
		return nil
	}
	pageCount, err := strconv.Atoi(userController.Ctx.QueryParam("pageCount"))
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warnln("User.GetPage query param pageCount must be a number.")
		userController.Ctx.ResStatus(http.StatusBadRequest)
		return nil
	}

	list, err := user.GetPage(page, pageCount)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Errorln("User.GetPage error")
		userController.Ctx.ResStatus(http.StatusInternalServerError)
		return nil
	}
	return list
}
