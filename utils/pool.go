package utils

import "sync"

type GoFuncPool struct {
	maxLimit  int
	gochannel chan struct{}
	sy        sync.WaitGroup
}

type PoolImpl interface {
	Submit(fn func())
	Close()
	Size() int
}

type GoFuncPoolOptions func(pool *GoFuncPool)

func WithMaxLimit(max int) GoFuncPoolOptions {
	return func(pool *GoFuncPool) {
		pool.maxLimit = max

		pool.gochannel = make(chan struct{}, pool.maxLimit)

		for i := 0; i < pool.maxLimit; i++ {
			pool.gochannel <- struct{}{}
		}
	}
}
func NewGoPool(options ...GoFuncPoolOptions) PoolImpl {
	pool := &GoFuncPool{}

	for _, option := range options {
		option(pool)
	}

	return pool
}

func (p *GoFuncPool) Submit(fn func()) {
	channel := <-p.gochannel
	p.sy.Add(1)
	go func() {
		fn()
		p.gochannel <- channel
		defer p.sy.Done()
	}()

	p.sy.Wait()
}

func (p *GoFuncPool) Close() {
	for i := 0; i < p.maxLimit; i++ {
		<-p.gochannel
	}

	close(p.gochannel)
}

func (p *GoFuncPool) Size() int {
	return len(p.gochannel)
}
