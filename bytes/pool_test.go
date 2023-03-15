package bytes_test

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"

	"github.com/powerpuffpenguin/easygo/algorithm"
	"github.com/powerpuffpenguin/easygo/bytes"
	"github.com/stretchr/testify/assert"
)

func TestPool(t *testing.T) {
	vals := []int{
		2, 4, 8, 16,
	}

	p := bytes.NewPool(algorithm.Map(vals, func(v int) bytes.BlockAllocator {
		return bytes.NewAllocatorPool(v, true, 10)
	}))

	for i, val := range vals {
		b := p.Get(val + 1)
		if !assert.Equal(t, len(b), val+1) {
			t.FailNow()
		}
		if i+1 == len(vals) {
			if !assert.Equal(t, val+1, cap(b)) {
				t.FailNow()
			}
		} else {
			if !assert.Equal(t, vals[i+1], cap(b)) {
				t.FailNow()
			}
		}
	}

	for _, val := range vals {
		p.Put(make([]byte, val+1))
	}
	for _, val := range vals {
		b := p.Get(val)
		if !assert.Equal(t, val, len(b)) {
			t.FailNow()
		}
		if !assert.Equal(t, val+1, cap(b)) {
			t.FailNow()
		}
	}
}

var minBlocks = []int{
	128,
	256,
	512,
	1024,
	128,
	256,
	512,
	1024,
}

func benchmarkPool(b *testing.B, p *bytes.Pool, vals []int) {
	runtime.GC()
	b.ResetTimer()

	var w0 sync.WaitGroup
	count := runtime.GOMAXPROCS(0)

	ch := make(chan []byte, count*4)
	for i := 0; i < count*2; i++ {
		w0.Add(1)
		go func() {
			for b := range ch {
				if p != nil {
					p.Put(b)
				}
			}
			w0.Done()
		}()
	}
	for i := 0; i < b.N; i++ {
		for _, val := range vals {
			size := rand.Intn(val/2) + val/2
			if p == nil {
				ch <- make([]byte, size)
			} else {
				ch <- p.Get(size)
			}
		}
	}
	close(ch)
	w0.Wait()
}
func newPool(sync bool, cache int) *bytes.Pool {
	return bytes.NewPool(
		[]bytes.BlockAllocator{
			bytes.NewAllocatorPool(128, sync, cache),
			bytes.NewAllocatorPool(256, sync, cache),
			bytes.NewAllocatorPool(512, sync, cache),
			bytes.NewAllocatorPool(1024, sync, cache),
			bytes.NewAllocatorPool(2048, sync, cache),
			bytes.NewAllocatorPool(4096, sync, cache),
			bytes.NewAllocatorPool(8192, sync, cache),
			bytes.NewAllocatorPool(16384, sync, cache),
			bytes.NewAllocatorPool(32768, sync, cache),
		},
	)
}

func BenchmarkPoolMinNo(b *testing.B) {
	benchmarkPool(b, nil, minBlocks)
}

func BenchmarkPoolMinSync(b *testing.B) {
	benchmarkPool(b, newPool(true, 0), minBlocks)
}

func BenchmarkPoolMinCache(b *testing.B) {
	benchmarkPool(b, newPool(false, 100), minBlocks)
}
func BenchmarkPoolMin(b *testing.B) {
	benchmarkPool(b, newPool(true, 100), minBlocks)
}

func BenchmarkPoolMin4No(b *testing.B) {
	benchmarkPool(b, nil, minBlocks[:4])
}

func BenchmarkPoolMin4Sync(b *testing.B) {
	benchmarkPool(b, newPool(true, 0), minBlocks[:4])
}

func BenchmarkPoolMin4Cache(b *testing.B) {
	benchmarkPool(b, newPool(false, 100), minBlocks[:4])
}
func BenchmarkPoolMin4(b *testing.B) {
	benchmarkPool(b, newPool(true, 100), minBlocks[:4])
}

var mediumBlocks = []int{
	2048,
	4096,
	8192,
	16384,
	32768,
	2048,
	4096,
	8192,
	16384,
	32768,
}

func BenchmarkPoolMediumNo(b *testing.B) {
	benchmarkPool(b, nil, mediumBlocks)
}

func BenchmarkPoolMediumSync(b *testing.B) {
	benchmarkPool(b, newPool(true, 0), mediumBlocks)
}

func BenchmarkPoolMediumCache(b *testing.B) {
	benchmarkPool(b, newPool(false, 100), mediumBlocks)
}
func BenchmarkPoolMedium(b *testing.B) {
	benchmarkPool(b, newPool(true, 100), mediumBlocks)
}

func BenchmarkPoolMedium5No(b *testing.B) {
	benchmarkPool(b, nil, mediumBlocks[:4])
}

func BenchmarkPoolMedium5Sync(b *testing.B) {
	benchmarkPool(b, newPool(true, 0), mediumBlocks[:4])
}

func BenchmarkPoolMedium5Cache(b *testing.B) {
	benchmarkPool(b, newPool(false, 100), mediumBlocks[:4])
}
func BenchmarkPoolMedium5(b *testing.B) {
	benchmarkPool(b, newPool(true, 100), mediumBlocks[:4])
}

var blocks = []int{
	128,
	256,
	512,
	1024,
	2048,
	4096,
	8192,
	16384,
	32768,
}

func BenchmarkPoolBlocksNo(b *testing.B) {
	benchmarkPool(b, nil, blocks)
}

func BenchmarkPoolBlocksSync(b *testing.B) {
	benchmarkPool(b, newPool(true, 0), blocks)
}

func BenchmarkPoolBlocksCache(b *testing.B) {
	benchmarkPool(b, newPool(false, 100), blocks)
}
func BenchmarkPoolBlocks(b *testing.B) {
	benchmarkPool(b, newPool(true, 100), blocks)
}
