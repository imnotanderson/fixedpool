package fixdpool

import "sync"

type FixedPool struct {
	list      []interface{}
	newFunc   func() interface{}
	resetFunc func(interface{}) interface{}
	lock      *sync.RWMutex
}

func (p *FixedPool) Get() interface{} {
	p.lock.Lock()
	defer p.lock.Unlock()
	l := len(p.list)
	if l > 0 {
		data := p.list[l-1]
		p.list[l-1] = nil
		p.list = p.list[:l-1]
		return data
	}
	return p.newFunc()
}

func (p *FixedPool) Put(obj interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()
	l := len(p.list)
	if l < cap(p.list) {
		p.list = append(p.list, p.resetFunc(obj))
	}
}

func NewFixedPool(size int, newFunc func() interface{}, resetFunc func(interface{}) interface{}) *FixedPool {
	return &FixedPool{
		list:      make([]interface{}, 0, size),
		newFunc:   newFunc,
		resetFunc: resetFunc,
		lock:      new(sync.RWMutex),
	}
}
