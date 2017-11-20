package conn

import (
	"container/list"
	"sync"
)

type Conn interface{}
type ConnPool interface {
	Get() Conn
	Put(Conn)
}

type pool struct {
	size int
	l    *list.List
	c    chan struct{}
	m    sync.Mutex
}

func (p *pool) init(size int) {
	if size == 0 {
		size = 100
	}
	p.size = size
	p.l = list.New()
	p.c = make(chan struct{}, size)
}
