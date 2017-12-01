package config

import (
	"time"

	"github.com/BurntSushi/toml"
)

var Default Config

type Config struct {
	Base    base
	MongoDB mongodb
	Redis   redis
}

// 基础配置
type base struct {
	Address string
	SlowRes time.Duration
}

// MongoDB 配置
type mongodb struct {
	Host         string
	Password     string
	DatebaseName string
	PoolSize     int
	SlowRes      time.Duration
}

// redis 配置
type redis struct {
	Host     string
	Password string
	PoolSize int
	SlowRes  time.Duration
}

// 创建一个Config对象
func New(fileName string) error {
	if _, err := toml.DecodeFile(fileName, &Default); err != nil {
		return err
	}
	Default.Base.SlowRes = Default.Base.SlowRes * 1000 * 1000
	Default.MongoDB.SlowRes = Default.MongoDB.SlowRes * 1000 * 1000
	Default.Redis.SlowRes = Default.Redis.SlowRes * 1000 * 1000
	return nil
}
