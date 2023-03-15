package bytes

import (
	"strconv"
	"sync"
)

// memory allocation release interface
type BlockAllocator interface {
	// Return block size
	Block() int
	// Memory is no longer in use release it
	Put(b []byte)
	// Get a block of memory
	Get() []byte
}

type AllocatorPool struct {
	blocksize int
	pool      *sync.Pool
	ch        chan []byte
}

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
	}
	return &AllocatorPool{
		blocksize: blocksize,
		pool:      p,
		ch:        ch,
	}
}

// Return block size
func (a *AllocatorPool) Block() int {
	return a.blocksize
}

// Memory is no longer in use release it
func (a *AllocatorPool) Put(b []byte) {
	if cap(b) < a.blocksize {
		return
	}
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

// Get a block of memory
func (a *AllocatorPool) Get() (b []byte) {
	if a.ch != nil {
		select {
		case b = <-a.ch:
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
