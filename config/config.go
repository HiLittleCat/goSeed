package config

import (
	"time"

	"github.com/BurntSushi/toml"
)

// Default 配置项变量
var (
	Default Config //默认配置变量
)

// Config 配置项信息
type Config struct {
	Base      base
	MongoDB   mongodb
	RedisBase redisBase
	Etcd      etcd
	Session   session
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
	Host     string
	Password string
	Name     string
	PoolSize int
	SlowRes  time.Duration
}

// redis 配置
type redisBase struct {
	Host     string
	Password string
	PoolSize int
	DB       int
	SlowRes  time.Duration
	Name     string
}

// etcd 配置
type etcd struct {
	UseEtcd     bool
	Endpoints   string
	DialTimeout time.Duration
}

// session 配置
type session struct {
	Name     string
	Expire   time.Duration
	Secure   bool
	HttpOnly bool
	Domain   string
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
	Default.RedisBase.SlowRes = Default.RedisBase.SlowRes * time.Millisecond
	Default.Session.Expire = Default.Session.Expire * time.Second
	if Default.Etcd.UseEtcd == false {
		return nil
	}
	return InitConifgFromEtcd()
}
