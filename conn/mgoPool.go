package conn

import (
	"gopkg.in/mgo.v2"
)

type MgoPool struct {
	p   pool
	opt MgoPoolOption
}

type MgoPoolOption struct {
	Size   int
	Host   string
	DbName string
}

func NewMgoPool(opt MgoPoolOption) (*MgoPool, error) {
	p := MgoPool{}
	err := p.init(opt)
	return &p, err
}

func (p *MgoPool) init(opt MgoPoolOption) error {
	p.opt = opt
	p.p.init(opt.Size)

	session, err := mgo.Dial(opt.Host)
	if err != nil {
		return err
	}
	for i := 0; i < p.p.size; i++ {
		p.p.c <- struct{}{}
	}
	p.p.l.PushBack(session)
	return nil
}

// 获取一个mongo连接
func (p *MgoPool) Get() Conn {
	_ = <-p.p.c
	p.p.m.Lock()
	defer p.p.m.Unlock()
	return p.p.l.Front().Value.(*mgo.Session).Clone()
}

// 释放一个mongo连接
func (p *MgoPool) Put(c Conn) {
	p.p.m.Lock()
	defer p.p.m.Unlock()
	c.(*mgo.Session).Close()
	p.p.c <- struct{}{}
}

// 使用连接池
func (p *MgoPool) Exec(collection string, callback func(*mgo.Collection)) {
	_session := p.Get().(*mgo.Session)
	defer func() {
		p.Put(_session)
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	c := _session.DB(p.opt.DbName).C(collection)
	callback(c)
}
