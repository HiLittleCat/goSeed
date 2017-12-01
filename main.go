package main

import (
	"os"
	"runtime"

	"github.com/HiLittleCat/goSeed/config"
	"github.com/HiLittleCat/goSeed/conn"
	"github.com/HiLittleCat/goSeed/controller"
	"github.com/HiLittleCat/goSeed/middleware"

	"github.com/HiLittleCat/core"

	log "github.com/sirupsen/logrus"
)

func main() {
	//本函数将在调度程序优化后会去掉。
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Parse config.ini
	if err := config.New("config/config.ini"); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	logFile, err := os.OpenFile("log/service.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0642)
	defer logFile.Close()
	log.SetOutput(logFile)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{})

	//mongodb init
	mgoOption := conn.MgoPoolOption{
		Host:   config.Default.MongoDB.Host,
		Size:   config.Default.MongoDB.PoolSize,
		DbName: config.Default.MongoDB.DatebaseName,
	}
	mgoPool, err := conn.NewMgoPool(mgoOption)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatalln("connect mongodb fail")
		os.Exit(1)
	}
	conn.MgoSet(conn.MgoBosh, mgoPool)

	// redis init
	redisOption := conn.RedisPoolOption{
		Host:     config.Default.Redis.Host,
		Password: config.Default.Redis.Password,
		Size:     config.Default.Redis.PoolSize,
	}
	redisPool, err := conn.NewRedisPool(redisOption)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatalln("connect redis fail")
		os.Exit(1)
	}
	conn.RedisSet(conn.RedisBosh, redisPool)

	// Middleware register
	core.Use(middleware.Container)
	core.Use(middleware.Session)

	// Controller register
	core.AutoController(&controller.User{})

	core.Address = config.Default.Base.Address
	// Run server
	core.Run()
}
