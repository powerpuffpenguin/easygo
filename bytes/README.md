# bytes

```
import "github.com/powerpuffpenguin/easygo/bytes"
```

# BytesToString/StringToBytes

BytesToString/StringToBytes functions are used to convert between strings and []byte, Faster than standard conversion because no reallocation of memory.

Because the string is stored in the specified memory, using StringToBytes to convert the original string to []byte is also read-only, writing will cause undefined behavior.

```
b := StringToBytes("123")
s := BytesToString(b)
```

# AllocatorPool

AllocatorPool can be seen as an upgraded version of sync.Pool for []byte.

AllocatorPool is used to allocate fixed size []byte, you can specify the cache size, and it is easy to automatically clean up the cache when idle

```
	allocator := bytes.NewAllocatorPool(128, // block size 128
		true, // enable sync.Pool
		100,  // chan cache 100
	)
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

	// allocation block
	b := allocator.Get()
	// use block
	fmt.Println(b)
	// Return blocks are no longer used
	allocator.Put(b)
```

# Pool

AllocatorPool is used to allocate fixed size []byte, and Pool is used to apply for arbitrary size []byte. It uses AllocatorPool at the bottom to apply for memory in blocks

```
    p := bytes.NewPool(
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
	// allocation []byte
	b := p.Get(128)
	// use []byte
	fmt.Println(b)
	// Return []byte are no longer used
	p.Put(b)
```

# Buffer

It is the same as the bytes.Buffer of the standard library, except that the Allocator property is used to specify where to apply for memory.