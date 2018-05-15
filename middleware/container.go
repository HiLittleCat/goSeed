package middleware

import (
	"time"

	"github.com/HiLittleCat/goSeed/config"

	"github.com/HiLittleCat/core"

	log "github.com/sirupsen/logrus"
)

// Container  打印访问日志
func Container(ctx *core.Context) {
	start := time.Now()
	ctx.Next()
	t := time.Since(start)
	if t >= config.Default.Base.SlowRes {
		log.Infoln(ctx.Request.Method, ctx.Request.URL, t)
	}
}
