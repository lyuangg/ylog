package ylog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceToStr(t *testing.T) {
	var a1 int = 1
	var a2 uint = 2
	var a3 uint8 = 3
	var a4 uint16 = 4
	var a5 int16 = 5
	var a6 uint32 = 6
	var a7 int32 = -7
	var a8 uint64 = 8
	var a9 int64 = 99999
	var a10 float32 = 1.22
	var a11 float64 = 5.22
	var a12 bool = false
	var a13 = []byte{228, 189, 160, 229, 165, 189}
	var a14 int8 = 88
	var tt = map[string]string{
		"good": "nice",
	}
	var f = Fields{
		"key": "val",
	}

	assert.Equal(t, InterfaceToStr(a1), "1")
	assert.Equal(t, InterfaceToStr(a2), "2")
	assert.Equal(t, InterfaceToStr(a3), "3")
	assert.Equal(t, InterfaceToStr(a4), "4")
	assert.Equal(t, InterfaceToStr(a5), "5")
	assert.Equal(t, InterfaceToStr(a6), "6")
	assert.Equal(t, InterfaceToStr(a7), "-7")
	assert.Equal(t, InterfaceToStr(a8), "8")
	assert.Equal(t, InterfaceToStr(a9), "99999")
	assert.Equal(t, InterfaceToStr(a10), "1.22")
	assert.Equal(t, InterfaceToStr(a11), "5.22")
	assert.Equal(t, InterfaceToStr(a12), "false")
	assert.Equal(t, InterfaceToStr(a13), "你好")
	assert.Equal(t, InterfaceToStr(a14), "88")
	assert.Equal(t, InterfaceToStr(tt), "map[good:nice]")
	assert.Equal(t, InterfaceToStr(f), "map[key:val]")
	assert.Equal(t, InterfaceToStr("strstr"), "strstr")
	assert.Equal(t, InterfaceToStr(struct{}{}), "{}")
}
