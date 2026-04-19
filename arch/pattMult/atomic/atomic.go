package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type A struct {
	i int64
	m sync.Mutex
}

func (a *A) Inc(delta int64) {
	// a.m.Lock()
	// defer a.m.Unlock()
	// a.i += delta

	atomic.AddInt64(&a.i, delta)
}

func (a *A) Get() int64 {
	// a.m.Lock()
	// defer a.m.Unlock()
	// return a.i
	return atomic.LoadInt64(&a.i)
}

type B struct {
}

func (b *B) II(c int) {
	for i := 0; i < c; i++ {
		a.Inc(1)
	}
	wg.Done()
}

type C struct {
}

func (c *C) DD(count int) {
	for i := 0; i < count; i++ {
		a.Inc(-1)
	}
	wg.Done()
}

var a A
var b B
var c C

var wg sync.WaitGroup

func main() {
	var nCount = 1000000
	a.i = 0

	wg.Add(1)
	go b.II(nCount)

	wg.Add(1)
	go c.DD(nCount)

	wg.Wait()
	fmt.Println("Sum", a.Get())
}
