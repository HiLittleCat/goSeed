package middleware

import (
	"net/http"
	"time"

	"github.com/HiLittleCat/goSeed/conn"

	"github.com/HiLittleCat/goSeed/config"

	"github.com/HiLittleCat/core"

	log "github.com/sirupsen/logrus"

	redis "gopkg.in/redis.v5"
)

func Container(ctx *core.Context) {
	start := time.Now()
	ctx.Next()
	t := time.Since(start)
	if t >= config.Default.Base.SlowRes {
		log.Infoln(ctx.Request.Method, ctx.Request.URL, t)
	}
}

func ResWrite(ctx *core.Context) {
	ctx.Next()
	if ctx.Written() == false {
		_, err := ctx.Success(http.StatusOK, ctx.ResData)
		if err != nil {
			panic(err)
		}
	}
}

//set session info
func Session(ctx *core.Context) {
	redisPool := conn.GetRedisPool(conn.RedisBosh)
	redisPool.Exec(conn.SessionDB, func(c *redis.Client) {
		// cmd := c.Get("session:")
		// if err := cmd.Err(); err != nil {
		// 	log.WithFields(log.Fields{"err": err}).Warn("get session fail")
		// }
		// cmd.Result()
		// ctx.Session = make(map[string]interface{})
		// ctx.Session["user"] = ctx.Request.UserAgent
	})
	ctx.Next()
}
