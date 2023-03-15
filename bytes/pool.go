package bytes

import (
	"sort"
	"strconv"
	"sync"
)

type bytesPool struct {
	size int
	pool sync.Pool
	ch   chan []byte
}

func (p *bytesPool) New() any {
	return make([]byte, p.size)
}
func (p *bytesPool) Get() (b []byte) {
	select {
	case b = <-p.ch:
	default:
		b = p.pool.Get().([]byte)
	}
	return
}
func (p *bytesPool) Put(b []byte) {
	select {
	case p.ch <- b:
	default:
		p.pool.Put(b)
	}
}

// Use the block memory implemented by sync.Pool.
// Reuse []byte to reduce the time of memory application and release.
type Pool struct {
	pools []*bytesPool
}

// Create a memory pool, size specifies the block size provided in the pool
func NewPool(size ...int) *Pool {
	var (
		pools []*bytesPool
		n     = len(size)
	)

	if n > 0 {
		sort.Ints(size)
		if size[0] < 1 {
			panic(`bytes.NewPool: capacity=` + strconv.Itoa(size[0]) + ` is invalid`)
		}

		pools = make([]*bytesPool, n)
		last := -1
		for i, v := range size {
			if v == last {
				panic(`bytes.NewPool: capacity=` + strconv.Itoa(v) + ` already exists`)
			}
			last = v

			pool := &bytesPool{
				size: v,
				ch:   make(chan []byte, 100),
			}
			pool.pool.New = pool.New
			pools[i] = pool
		}
	}
	return &Pool{
		pools: pools,
	}
}

// Like make([]byte,size,auto) from pool
func (p *Pool) Get(size int) (b []byte) {
	n := len(p.pools)
	if n == 0 {
		b = make([]byte, size)
		return
	}
	i := sort.Search(n, func(i int) bool { return p.pools[i].size >= size })
	if i == n { // not match
		b = make([]byte, size)
		return
	}
	b = p.pools[i].Get()
	b = b[:size]
	return
}

// Return bytes to the pool for reuse.
// Bytes may not be created by the pool, as long as it is confirmed that no one will use it, it can be submitted to the pool
func (p *Pool) Put(b []byte) {
	n := len(p.pools)
	if n == 0 {
		return
	}
	size := cap(b)
	if size == 0 {
		return
	}
	i := sort.Search(n, func(i int) bool { return p.pools[i].size >= size })
	if i == n { // not match
		i--
		if size <= p.pools[i].size*4/3 {
			p.pools[i].Put(b)
		}
		return
	}
	if p.pools[i].size == size { // equal
		p.pools[i].Put(b)
	} else if i > 0 {
		p.pools[i-1].Put(b)
	} // else not match
}
