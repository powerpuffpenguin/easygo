package easygo_test

import (
	"testing"

	"github.com/powerpuffpenguin/easygo"
	"github.com/stretchr/testify/assert"
)

func TestPair(t *testing.T) {
	v := easygo.Pair[int, string]{1, "ok"}
	if !assert.Equal(t, v.First, 1) {
		t.FailNow()
	}
	if !assert.Equal(t, v.Second, "ok") {
		t.FailNow()
	}

	v1 := easygo.MakePair(uint32(1), "ok")
	if !assert.Equal(t, v1.First, uint32(1)) {
		t.FailNow()
	}
	if !assert.Equal(t, v1.Second, "ok") {
		t.FailNow()
	}

	v2 := easygo.NewPair(uint32(1), "ok")
	if !assert.Equal(t, v2.First, uint32(1)) {
		t.FailNow()
	}
	if !assert.Equal(t, v2.Second, "ok") {
		t.FailNow()
	}
}
