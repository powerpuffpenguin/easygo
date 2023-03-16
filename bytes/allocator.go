package bytes

import (
	"strconv"
	"sync"
)

// This is a memory allocation and release interface.
// Although go comes with a memory pool,
// a custom memory pool can improve program efficiency and reduce fragmentation at some point.
type Allocator interface {
	// Memory is no longer in use release it
	Put(b []byte)
	// Get a block of memory
	Get(size int) []byte
}

// This is a memory block allocation and release interface.
// It is usually more efficient to release memory by block application than random size application release.
// In addition, chunks are easier to cache and reuse than random sizes.
type BlockAllocator interface {
	// Memory is no longer in use release it
	Put(b []byte)
	// Get a block of memory
	Get() []byte
	// Return block size
	Block() int
}

// This is a memory block allocator for reusing block memory
type AllocatorPool struct {
	blocksize, cachesize int
	pool                 *sync.Pool
	ch                   chan []byte
}

// Create a block memory allocator
//
// blocksize specifies the minimum size of a memory block,
// if pool is true the sync.Pool cache block will be used,
// if cachesize is greater than 0, cache blocks using make(chan []byte,cachesize),
func NewAllocatorPool(blocksize int,
	pool bool,
	cachesize int,
) *AllocatorPool {
	if blocksize < 0 {
		panic(`blocksize(` + strconv.Itoa(blocksize) + `) must >=0`)
	}
	var p *sync.Pool
	if pool {
		p = &sync.Pool{
			New: func() any {
				return make([]byte, blocksize)
			},
		}
	}
	var ch chan []byte
	if cachesize > 0 {
		ch = make(chan []byte, cachesize)
	} else {
		cachesize = 0
	}
	return &AllocatorPool{
		blocksize: blocksize,
		pool:      p,
		ch:        ch,
		cachesize: cachesize,
	}
}

// Return block size
func (a *AllocatorPool) Block() int {
	return a.blocksize
}

// Memory is no longer in use release it
func (a *AllocatorPool) Put(b []byte) {
	if cap(b) >= a.blocksize {
		if a.ch != nil {
			select {
			case a.ch <- b:
				return
			default:
			}
		}
		if a.pool != nil {
			a.pool.Put(b)
		}
	}
}

// Put b into sync.Pool if it exists
func (a *AllocatorPool) PutPool(b []byte) {
	if cap(b) >= a.blocksize {
		if a.pool != nil {
			a.pool.Put(b)
		}
	}
}

// Get a block of memory
func (a *AllocatorPool) Get() (b []byte) {
	if a.ch != nil {
		select {
		case b = <-a.ch:
			b = b[:a.blocksize]
			return
		default:
		}
	}
	if a.pool != nil {
		b = a.pool.Get().([]byte)
		return
	}
	return make([]byte, a.blocksize)
}

// Returns cache size
func (a *AllocatorPool) CacheSize() int {
	return a.cachesize
}

// Returns the cache chan, which can usually be used to define to clean up unused caches
func (a *AllocatorPool) Cache() chan []byte {
	return a.ch
}
