package middleware

import (
	"github.com/HiLittleCat/core"
	"github.com/HiLittleCat/goSeed/errors"
)

// RequireSession  需要登录的接口调用此中间件
func RequireSession(ctx *core.Context) {
	sses := ctx.GetSession()
	if sses == nil {
		ctx.Fail(errors.ErrUserExist)
		return
	}
	ctx.Next()
}
