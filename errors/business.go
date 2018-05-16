package errors

import (
	"github.com/HiLittleCat/core"
)

var businessErr = &core.BusinessError{}
var ErrUserExist = businessErr.New(10000, "user has exist")
var ErrNeedLogin = businessErr.New(10001, "need login")
