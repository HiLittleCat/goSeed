package middleware

import (
	"net/http"

	"github.com/HiLittleCat/core"
	"github.com/HiLittleCat/goSeed/session"
)

const cookieName = "PERFECT_FLY_ID"

// Session session处理
func Session(ctx *core.Context) {

	var cookie *http.Cookie
	cookies := ctx.Request.Cookies()
	if len(cookies) > 0 {
		cookie = cookies[0]
	}
	sessionID := cookie.Value
	sessionInfo, err := session.StoreS.Get(sessionID)
	if err != nil {
		ctx.Fail(err)
		return
	}
	if sessionInfo == nil || len(sessionInfo) == 0 {

	} else {
		ctx.Data["session"] = sessionInfo
		uidCookie := &http.Cookie{
			Name:     cookieName,
			Value:    sessionID,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   session.Expire,
		}
		http.SetCookie(ctx.ResponseWriter, uidCookie)
	}

	ctx.Next()
}
