package bytes_test

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/powerpuffpenguin/easygo/bytes"
	"github.com/stretchr/testify/assert"
)

func TestPool(t *testing.T) {
	vals := []int{
		2, 4, 8, 16,
	}
	p := bytes.NewPool(vals...)
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

func benchmarkPool(p *bytes.Pool, vals []int) {
	if p == nil {
		for _, val := range vals {
			_ = make([]byte, val)
		}
	} else {
		for _, val := range vals {
			b := p.Get(val)
			p.Put(b)
		}
	}
}

var blocks = []int{
	128,
	256,
	// 512,
	// 1024,
	// 2048,
	// 4096,
	// 8192,
	// 16384,
	// 32768,
}

func BenchmarkNoPool(b *testing.B) {
	vals := blocks

	var w0 sync.WaitGroup
	ch := make(chan []byte, 100)
	for i := 0; i < 100; i++ {
		w0.Add(1)
		go func() {
			for range ch {
			}
			w0.Done()
		}()
	}
	for i := 0; i < b.N; i++ {
		ch <- make([]byte, rand.Intn(128))
		for _, val := range vals {
			ch <- make([]byte, rand.Intn(val))
		}
	}
	close(ch)
	w0.Wait()
}
func BenchmarkPool(b *testing.B) {
	vals := blocks
	p := bytes.NewPool(vals...)

	var w0 sync.WaitGroup
	ch := make(chan []byte, 100)
	for i := 0; i < 100; i++ {
		w0.Add(1)
		go func() {
			for b := range ch {
				p.Put(b)
			}
			w0.Done()
		}()
	}
	for i := 0; i < b.N; i++ {
		ch <- p.Get(rand.Intn(128))
		for _, val := range vals {
			ch <- p.Get(rand.Intn(val))
		}
	}
	close(ch)
	w0.Wait()
}
