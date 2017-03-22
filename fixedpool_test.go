package fixdpool

import "testing"

func TestNewFixedPool(t *testing.T) {
	type A struct {
		i int
	}
	p := NewFixedPool(1,
		func() interface{} { return &A{} },
		func(a interface{}) interface{} {
			c := a.(*A)
			c.i = 0
			return c
		})
	d := p.Get()
	d.(*A).i = 1
	p.Put(d)
	if len(p.list) != 1 {
		t.Error("put err:", len(p.list) != 1)
	}
	if d.(*A).i != 0 {
		t.Error("reset err:", d.(*A).i)
	}
	d.(*A).i = 2
	p.Put(d)
	if d.(*A).i != 2 {
		t.Error("put err,d reset")
	}
	if len(p.list) != 1 {
		t.Error("put err,d reset")
	}
	p.Get()
	p.Put(d)
	if d.(*A).i != 0 {
		t.Error("put err")
	}
}
