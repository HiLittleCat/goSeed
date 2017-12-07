package controller

import (
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

func (ctr *User) Get() *model.User {
	user := model.User{Id: ctr.Ctx.QueryParam("_id")}
	if err := user.Get(); err != nil {
		log.WithFields(log.Fields{"err": err}).Warnln("UserController.Get error")
		ctr.Ctx.ResStatus(http.StatusInternalServerError)
		return nil
	}
	return &user
}

/**
* User/All
 */
func (ctr *User) GetAll() []model.User {
	user := model.User{}
	list, err := user.GetAll()
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warnln("UserController.GetAll error")
		ctr.Ctx.ResStatus(http.StatusInternalServerError)
		return nil
	}
	return list
}
