package main

import (
	"os"
	"runtime"

	"github.com/HiLittleCat/compress"
	"github.com/HiLittleCat/core"
	"github.com/HiLittleCat/goSeed/config"
	"github.com/HiLittleCat/goSeed/conn"
	"github.com/HiLittleCat/goSeed/controller"
	"github.com/HiLittleCat/goSeed/middleware"

	logcore "github.com/HiLittleCat/log"
	log "github.com/sirupsen/logrus"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	//本函数将在调度程序优化后会去掉。
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Parse config.ini
	if err := config.New("config/config.ini"); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Log config
	if b, _ := PathExists("log"); b == false {
		os.Mkdir("log", 0777)
	}
	logFile, err := os.OpenFile("log/service.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0642)
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetLevel(log.WarnLevel)
	log.SetFormatter(&log.TextFormatter{})

	// Mongodb init
	mgoOption := conn.MgoPoolOption{
		Host:   config.Default.MongoDB.Host,
		Size:   config.Default.MongoDB.PoolSize,
		DbName: config.Default.MongoDB.DatebaseName,
	}
	mgoPool, err := conn.NewMgoPool(mgoOption)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatalln("connect mongodb: " + config.Default.MongoDB.DatebaseName + "  fail")
		os.Exit(1)
	}
	conn.MgoSet(conn.MgoBosh, mgoPool)

	// Redis init
	redisOption := conn.RedisPoolOption{
		Host:     config.Default.Redis.Host,
		Password: config.Default.Redis.Password,
		Size:     config.Default.Redis.PoolSize,
	}
	redisPool, err := conn.NewRedisPool(redisOption)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatalln("connect redis:" + config.Default.Redis.Host + " fail")
		os.Exit(1)
	}
	conn.RedisSet(conn.RedisBosh, redisPool)

	// Middleware register
	logcore.Use()
	core.Use(middleware.Container)
	compress.Use()
	core.Use(middleware.ResWrite)
	core.Use(middleware.Session)

	// Controller register
	core.AutoController(&controller.User{})

	core.Address = config.Default.Base.Address
	// Run server
	core.Run()
}
