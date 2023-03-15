package bytes_test

import (
	"testing"
	"unsafe"

	"github.com/powerpuffpenguin/easygo/bytes"
	"github.com/stretchr/testify/assert"
)

func TestBytesToString(t *testing.T) {
	s0 := "01234567890"
	b0 := []byte(s0)
	if !assert.NotEqual(t, *(*int)(unsafe.Pointer(&s0)), *(*int)(unsafe.Pointer(&b0))) {
		t.FailNow()
	}

	s1 := bytes.BytesToString(b0)
	if !assert.Equal(t, *(*int)(unsafe.Pointer(&b0)), *(*int)(unsafe.Pointer(&s1))) {
		t.FailNow()
	}

	b1 := bytes.StringToBytes(s1)
	if !assert.Equal(t, *(*int)(unsafe.Pointer(&b0)), *(*int)(unsafe.Pointer(&b1))) {
		t.FailNow()
	}

	b2 := string(s1)
	if !assert.NotEmpty(t, *(*int)(unsafe.Pointer(&b0)), *(*int)(unsafe.Pointer(&b2))) {
		t.FailNow()
	}
}
