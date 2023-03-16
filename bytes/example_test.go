package bytes_test

import (
	"time"

	"github.com/powerpuffpenguin/easygo/bytes"
)

func ExamplePool() {
	p := bytes.NewPool(
		[]bytes.BlockAllocator{
			bytes.NewAllocatorPool(128, true, 100),
			bytes.NewAllocatorPool(1024, true, 100),
		})
	b0 := p.Get(128)
	b1 := p.Get(1024)
	copy(b0, b1)
	p.Put(b0)
	p.Put(b1)
}

func ExampeAllocatorPool_Cache() {
	allocator := bytes.NewAllocatorPool(128, true, 100)
	ch := allocator.Cache()
	if ch != nil {
		go func() {
			// Clean up 1/10 each time, so that a large number of caches will not be invalidated during peak periods,
			// and unused caches can be slowly cleaned up when the cache is not in use
			count := (allocator.CacheSize() + 9) / 10
			for {
				// Clear the cache every five minutes
				time.Sleep(time.Minute * 5)

			CF:
				for i := 0; i < count; i++ {
					select {
					case b := <-ch:
						allocator.PutPool(b)
					default:
						break CF
					}
				}
			}
		}()
	}
}
