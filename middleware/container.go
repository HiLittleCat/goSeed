package middleware

import (
	"time"

	"github.com/HiLittleCat/goSeed/conn"

	"github.com/HiLittleCat/goSeed/config"

	"github.com/HiLittleCat/core"

	log "github.com/sirupsen/logrus"

	redis "gopkg.in/redis.v5"
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

// Session session处理
func Session(ctx *core.Context) {
	redisPool := conn.GetRedisPool(conn.RedisBosh)
	redisPool.Exec(conn.SessionDB, func(c *redis.Client) {
		// cmd := c.Get("session:")
		// if err := cmd.Err(); err != nil {
		// 	log.WithFields(log.Fields{"err": err}).Warn("get session fail")
		// }
		// cmd.Result()
	})
	ctx.Next()
}
