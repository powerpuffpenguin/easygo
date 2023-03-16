package bytes

import (
	"sort"
	"strconv"
)

// Use the block memory implemented by sync.Pool.
// Reuse []byte to reduce the time of memory application and release.
//
// For block allocations smaller than 1024 the pool may be slower than the original make
// because the time spent looking for the allocator may exceed the memory allocation time
type Pool struct {
	opts       poolOptions
	allocators []BlockAllocator
}

// Create a memory pool, size specifies the block size provided in the pool
func NewPool(allocators []BlockAllocator, opt ...PoolOption) *Pool {
	if len(allocators) > 0 {
		sort.Slice(allocators, func(i, j int) bool {
			return allocators[i].Block() < allocators[j].Block()
		})
		var (
			min  = -1
			size int
		)
		for _, alloc := range allocators {
			size = alloc.Block()
			if min < size {
				min = size
			} else {
				panic(`easygo/bytes.NewPool Block(` + strconv.Itoa(size) + `) repeat`)
			}
		}
	}
	p := &Pool{
		opts:       defaultPoolOptions,
		allocators: allocators,
	}
	for _, o := range opt {
		o.Apply(&p.opts)
	}
	return p
}

// Like make([]byte,size,auto) from pool
func (p *Pool) Get(size int) (b []byte) {
	if p.opts.beforeGetf != nil {
		b = p.opts.beforeGetf(size)
		if b != nil {
			return
		}
	} else if size == 0 {
		return make([]byte, 0)
	}

	n := len(p.allocators)
	if n == 0 {
		if p.opts.newf == nil {
			b = make([]byte, size)
		} else {
			b = p.opts.newf(size)
		}
		return
	}
	i := sort.Search(n, func(i int) bool { return p.allocators[i].Block() >= size })
	if i == n { // not match
		if p.opts.newf == nil {
			b = make([]byte, size)
		} else {
			b = p.opts.newf(size)
		}
		return
	}
	if p.opts.getf == nil {
		b = p.allocators[i].Get()[:size]
	} else {
		b = p.opts.getf(p.allocators[i], size)
	}
	return
}

// Return bytes to the pool for reuse.
// Bytes may not be created by the pool, as long as it is confirmed that no one will use it, it can be submitted to the pool
func (p *Pool) Put(b []byte) {
	if p.opts.beforePutf != nil && p.opts.beforePutf(b) {
		return
	}

	n := len(p.allocators)
	if n == 0 {
		return
	}
	size := cap(b)
	if size == 0 {
		return
	}
	i := sort.Search(n, func(i int) bool { return p.allocators[i].Block() >= size })
	if i == n { // mismatch too large
		i--
		if p.opts.putf == nil {
			if size <= p.allocators[i].Block()*4/3 {
				p.allocators[i].Put(b)
			}
		} else {
			p.opts.putf(p.allocators[i], b)
		}
		return
	}
	if size == p.allocators[i].Block() { // equal
		if p.opts.putf == nil {
			p.allocators[i].Put(b)
		} else {
			p.opts.putf(p.allocators[i], b)
		}
	} else if i > 0 {
		i--
		if p.opts.putf == nil {
			if size <= p.allocators[i].Block()*4/3 {
				p.allocators[i].Put(b)
			}
		} else {
			p.opts.putf(p.allocators[i], b)
		}
	} // mismatch too small
}
