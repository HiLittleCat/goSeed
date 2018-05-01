package errors

import (
	"github.com/HiLittleCat/core"
)

var businessErr = &core.BusinessError{}
var ERR_USER_EXIST = businessErr.New(10000, "user has exist")
