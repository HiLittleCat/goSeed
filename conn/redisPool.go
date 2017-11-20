package conn

import (
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
}

func NewRedisPool(opt RedisPoolOption) (*RedisPool, error) {
	var p RedisPool
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
	p.p.c <- struct{}{}
}

// 使用连接池
func (self *RedisPool) Exec(callback func(*redis.Client)) {
	client := self.Get().(*redis.Client)
	defer func() {
		self.Put(client)
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	callback(client)
}
