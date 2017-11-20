package middleware

import (
	"common/conn"
	"core"
	"time"

	log "github.com/sirupsen/logrus"

	redis "gopkg.in/redis.v5"
)

func Container(ctx *core.Context) {
	start := time.Now()
	ctx.Next()
	if ctx.Written() == false {
		_, err := ctx.JSON(ctx.ResData)
		if err != nil {
			panic(err)
		}
	}
	log.Infoln(" %s  %s  %s", ctx.Request.Method, ctx.Request.URL, time.Since(start))
}

//set session info
func Session(ctx *core.Context) {
	redisPool := conn.GetRedisPool(conn.RedisBosh)
	redisPool.Exec(func(c *redis.Client) {
		cmd := c.Get("session:")
		if err := cmd.Err(); err != nil {
			//TODO LOG
		}
		//ctx.Session = make(map[string]interface{})
		//ctx.Session["UserAgent"] = ctx.Request.UserAgent
	})
	ctx.Next()
}
