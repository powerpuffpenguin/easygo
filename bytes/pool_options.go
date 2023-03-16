package bytes

import (
	"github.com/powerpuffpenguin/easygo/option"
)

type PoolOption option.Option[poolOptions]

var defaultPoolOptions = poolOptions{}

type poolOptions struct {
	beforeGetf func(size int) []byte
	beforePutf func(b []byte) bool
	newf       func(size int) []byte
	getf       func(allocator BlockAllocator, size int) []byte
	putf       func(allocator BlockAllocator, b []byte)
}

func PoolBeforeGet(f func(size int) []byte) PoolOption {
	return option.New(func(opts *poolOptions) {
		opts.beforeGetf = f
	})
}
func PoolBeforePut(f func(b []byte) bool) PoolOption {
	return option.New(func(opts *poolOptions) {
		opts.beforePutf = f
	})
}

// Call this function to allocate memory when there is no matching allocator,
// or return make([]byte, size) if the function f is nil
func PoolNew(f func(size int) []byte) PoolOption {
	return option.New(func(opts *poolOptions) {
		opts.newf = f
	})
}

// Call this function to allocate memory when a matching allocator is found,
// or return allocator.Get()[:size] if the function f is nil.
func PoolGet(f func(allocator BlockAllocator, size int) []byte) PoolOption {
	return option.New(func(opts *poolOptions) {
		opts.getf = f
	})
}

func PoolPut(f func(allocator BlockAllocator, b []byte)) PoolOption {
	return option.New(func(opts *poolOptions) {
		opts.putf = f
	})
}
