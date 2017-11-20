package config

import "github.com/BurntSushi/toml"

var Default Config

type Config struct {
	Base    base
	MongoDB mongodb
	Redis   redis
}

// 基础配置
type base struct {
	Address string
}

// MongoDB 配置
type mongodb struct {
	Host         string
	Password     string
	DatebaseName string
	PoolSize     int
}

// redis 配置
type redis struct {
	Host     string
	Password string
	DB       int
	PoolSize int
}

// 创建一个Config对象
func New(fileName string) error {
	if _, err := toml.DecodeFile(fileName, &Default); err != nil {
		return err
	}
	return nil
}
