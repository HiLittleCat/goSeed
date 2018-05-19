package main

import (
	"net/http"
	"os"
	"runtime"

	"github.com/HiLittleCat/compress"
	"github.com/HiLittleCat/conn"
	"github.com/HiLittleCat/core"
	logcore "github.com/HiLittleCat/log"
	"github.com/HiLittleCat/session"

	log "github.com/sirupsen/logrus"

	"github.com/HiLittleCat/goSeed/config"
	"github.com/HiLittleCat/goSeed/lib"
	"github.com/HiLittleCat/goSeed/middleware"
	_ "github.com/HiLittleCat/goSeed/routers"
)

func main() {

	// 本函数将在调度程序优化后会去掉。
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Parse config.ini
	if err := config.New("config/config.ini"); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Log config
	if b, _ := lib.PathExists("log"); b == false {
		os.Mkdir("log", 0777)
	}
	logFile, _ := os.OpenFile("log/service.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0642)
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetLevel(log.WarnLevel)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:10:10",
	})

	// Mongodb init
	mgoOption := conn.MgoPoolOption{
		Host:   config.Default.MongoDB.Host,
		Size:   config.Default.MongoDB.PoolSize,
		DbName: config.Default.MongoDB.Name,
	}
	mgoPool, err := conn.NewMgoPool(mgoOption)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatalln("connect mongodb: " + config.Default.MongoDB.Name + "  fail")
		os.Exit(1)
	}
	conn.MgoSet(mgoOption.DbName, mgoPool)

	// RedisBase init
	redisOption := conn.RedisPoolOption{
		Host:     config.Default.RedisBase.Host,
		Password: config.Default.RedisBase.Password,
		Size:     config.Default.RedisBase.PoolSize,
		DB:       config.Default.RedisBase.DB,
		SlowRes:  config.Default.RedisBase.SlowRes,
	}
	redisPool, err := conn.NewRedisPool(redisOption)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatalln("connect RedisBase:" + config.Default.RedisBase.Host + " fail")
		os.Exit(1)
	}
	conn.RedisSet(config.Default.RedisBase.Name, redisPool)

	// RedisSession init
	redisSessionOption := conn.RedisPoolOption{
		Host:     config.Default.RedisSession.Host,
		Password: config.Default.RedisSession.Password,
		Size:     config.Default.RedisSession.PoolSize,
		DB:       config.Default.RedisSession.DB,
		SlowRes:  config.Default.RedisSession.SlowRes,
	}
	redisSessionPool, err := conn.NewRedisPool(redisSessionOption)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatalln("connect RedisSession:" + config.Default.RedisSession.Host + " fail")
		os.Exit(1)
	}
	conn.RedisSet(config.Default.RedisSession.Name, redisSessionPool)

	// Core set
	core.Address = config.Default.Base.Address
	core.WriteTimeout = config.Default.Base.WriteTimeout
	core.ReadTimeout = config.Default.Base.ReadTimeout
	core.ListenLimit = config.Default.Base.ListenLimit
	core.Production = config.Default.Base.Production

	// Middleware register
	logcore.Use()
	core.Use(middleware.Container)

	session.Use(
		config.Default.Session.Expire,
		redisSessionPool,
		http.Cookie{
			Name:     config.Default.Session.Name,
			HttpOnly: config.Default.Session.HttpOnly,
			Domain:   config.Default.Session.Domain,
			Secure:   config.Default.Session.Secure,
		},
	)

	compress.Use()

	// Run server
	core.Run()
}
