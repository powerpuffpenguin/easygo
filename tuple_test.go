package easygo_test

import (
	"testing"

	"github.com/powerpuffpenguin/easygo"
	"github.com/stretchr/testify/assert"
)

func TestTuple(t *testing.T) {
	v := easygo.Tuple2[int, string]{1, "ok"}
	if !assert.Equal(t, v.V1, 1) {
		t.FailNow()
	}
	if !assert.Equal(t, v.V2, "ok") {
		t.FailNow()
	}

	v1 := easygo.MakeTuple2(uint32(1), "ok")
	if !assert.Equal(t, v1.V1, uint32(1)) {
		t.FailNow()
	}
	if !assert.Equal(t, v1.V2, "ok") {
		t.FailNow()
	}

	v2 := easygo.NewTuple2(uint32(1), "ok")
	if !assert.Equal(t, v2.V1, uint32(1)) {
		t.FailNow()
	}
	if !assert.Equal(t, v2.V2, "ok") {
		t.FailNow()
	}
}
