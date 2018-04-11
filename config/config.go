package config

import (
	"time"

	"github.com/BurntSushi/toml"
)

// Default 配置项变量
var (
	Default Config  //默认配置变量
	UseEtcd = false //是否通过etcd统一配置
)

// Config 配置项信息
type Config struct {
	Base    base
	MongoDB mongodb
	Redis   redis
	Etcd    etcd
}

// base 基础配置
type base struct {
	Address      string
	SlowRes      time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	ListenLimit  int
	Production   bool
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

// etcd 配置
type etcd struct {
	Endpoints   string
	DialTimeout time.Duration
}

// New 创建一个Config对象
func New(fileName string) error {
	if _, err := toml.DecodeFile(fileName, &Default); err != nil {
		return err
	}
	Default.Base.SlowRes = Default.Base.SlowRes * time.Millisecond
	Default.Base.WriteTimeout = Default.Base.WriteTimeout * time.Millisecond
	Default.Base.ReadTimeout = Default.Base.ReadTimeout * time.Millisecond
	Default.MongoDB.SlowRes = Default.MongoDB.SlowRes * time.Millisecond
	Default.Redis.SlowRes = Default.Redis.SlowRes * time.Millisecond
	InitEtcdConifg()
	return nil
}
