package conn

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/redis.v5"
)

type RedisPool struct {
	p   pool
	opt RedisPoolOption
}

type RedisPoolOption struct {
	Size     int
	Host     string
	Password string
	DB       int
	SlowRes  time.Duration
}

func NewRedisPool(opt RedisPoolOption) (*RedisPool, error) {
	var p RedisPool
	if opt.SlowRes == 0 {
		opt.SlowRes = time.Millisecond * 100
	}
	err := p.init(opt)
	return &p, err
}

func (p *RedisPool) init(opt RedisPoolOption) error {
	p.opt = opt
	p.p.init(opt.Size)

	redisOpt := redis.Options{
		Addr:     opt.Host,
		Password: opt.Password,
		PoolSize: opt.Size,
		DB:       opt.DB,
	}
	c := redis.NewClient(&redisOpt)
	for i := 0; i < p.p.size; i++ {
		p.p.c <- struct{}{}
	}
	p.p.l.PushBack(c)
	return nil
}

func (p *RedisPool) Close() {
	p.p.m.Lock()
	defer p.p.m.Unlock()

	client := p.p.l.Front()
	client.Value.(*redis.Client).Close()
}

// 获取一个redis连接
func (p *RedisPool) Get() Conn {
	_ = <-p.p.c
	p.p.m.Lock()
	defer p.p.m.Unlock()
	return p.p.l.Front().Value.(*redis.Client)
}

// 释放一个redis连接
func (p *RedisPool) Put(c Conn) {
	p.p.m.Lock()
	defer p.p.m.Unlock()
	c.(*redis.Client).Close()
	p.p.c <- struct{}{}
}

// 使用连接池
func (p *RedisPool) Exec(callback func(*redis.Client)) {
	start := time.Now()
	client := p.Get().(*redis.Client)
	defer func() {
		p.Put(client)
		if err := recover(); err != nil {
			log.Errorln("redis exec err, ", err)
			panic(err)
		}
		t := time.Since(start)
		if t >= p.opt.SlowRes && p.opt.SlowRes != 0 {
			log.Warnln("redis exec db:", p.opt.DB, t)
		}
	}()
	callback(client)
}
