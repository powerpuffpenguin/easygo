package bytes

import (
	"github.com/powerpuffpenguin/easygo/option"
)

type PoolOption option.Option[poolOptions]

var defaultPoolOptions = poolOptions{}

type poolOptions struct {
	newf func(size int) []byte
	getf func(allocator BlockAllocator, size int) []byte
	// putf
}

// Call this function to allocate memory when there is no matching allocator,
// or return make([]byte, size) if the function is nil
func PoolNew(f func(size int) []byte) PoolOption {
	return option.New(func(opts *poolOptions) {
		opts.newf = f
	})
}

// Call this function to allocate memory when a matching allocator is found,
// or return allocator.Get(size) if the function is nil
func PoolGet(f func(allocator BlockAllocator, size int) []byte) PoolOption {
	return option.New(func(opts *poolOptions) {
		opts.getf = f
	})
}
