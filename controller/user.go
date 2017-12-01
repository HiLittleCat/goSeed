package controller

import (
	"github.com/HiLittleCat/goSeed/model"

	"github.com/HiLittleCat/core"

	"net/http"

	log "github.com/sirupsen/logrus"
)

type User struct {
	core.Controller
}

func (this *User) Get() interface{} {
	user := model.User{Id: this.Ctx.QueryParam("_id")}
	if err := user.Get(); err != nil {
		log.WithFields(log.Fields{"err": err}).Warnln("UserController.Get error")
		this.Ctx.ResStatus(http.StatusInternalServerError)
		return nil
	}

	return user
}
