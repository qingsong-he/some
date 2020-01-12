package main

import (
	. "github.com/qingsong-he/ce"
	"sync"
	"unsafe"
)

func main() {
	p := &sync.Pool{
		New: func() interface{} {
			return &sync.Map{}
		},
	}

	m1 := p.Get().(*sync.Map)
	m2 := p.Get().(*sync.Map)
	m1.Store(1, nil)
	m2.Store(3, nil)
	p.Put(m1)
	p.Put(m2)

	m11 := p.Get().(*sync.Map)
	m22 := p.Get().(*sync.Map)
	Print(unsafe.Pointer(m1) == unsafe.Pointer(m11)) // true
	Print(unsafe.Pointer(m2) == unsafe.Pointer(m22)) // true

	m11.Range(func(key, value interface{}) bool {
		Print(key.(int)) // 1
		return true
	})
	m22.Range(func(key, value interface{}) bool {
		Print(key.(int)) // 3
		return true
	})
}
