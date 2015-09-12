package bitarr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const bit1 byte = 0x1
const bit0 byte = 0x0

// Test if creation and lengths are ok
func TestBitarrNew(t *testing.T) {

	arr := NewArray(64)
	assert.NotNil(t, arr)

	assert.Equal(t, uint32(64), arr.Len(), "Length is wrong")
	assert.Equal(t, 8, len(arr.b), "Too much/less memory allocated")

	arr = NewArray(65)
	assert.Equal(t, uint32(65), arr.Len(), "Length is wrong")
	assert.Equal(t, 9, len(arr.b), "Too much/less memory allocated")

	for i := uint32(0); i < arr.Len(); i++ {
		assert.Equal(t, bit0, arr.Get(i), "Not correctly initialised")
	}
}

func TestBitarrPanics(t *testing.T) {

	arr := NewArray(64)

	assert.Panics(t, func() {
		arr.Get(65)
	}, "Should panic with an out-of-bounds")

	assert.Panics(t, func() {
		arr.Get(1337)
	}, "Should panic with an out-of-bounds")
}

// Test if get, set and clear work as expected
func TestGetSet(t *testing.T) {

	arr := NewArray(65)

	arr.Set(33, 1)
	arr.Set(23, 1)
	arr.Set(18, 1)
	arr.Set(17, 0)

	assert.Equal(t, bit1, arr.Get(33), "False negative")
	assert.Equal(t, bit1, arr.Get(23), "False negative")
	assert.Equal(t, bit1, arr.Get(18), "False negative")
	assert.Equal(t, bit0, arr.Get(17), "False negative")
	assert.Equal(t, bit0, arr.Get(15), "False negative")

	arr.Set(33, 0)
	arr.Set(18, 0)

	assert.Equal(t, bit0, arr.Get(33), "False positive")
	assert.Equal(t, bit1, arr.Get(23), "False positive")
	assert.Equal(t, bit0, arr.Get(18), "False positive")

}
