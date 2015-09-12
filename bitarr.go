// Package bitarr provides a bitarray, can be helpful to efficiently store array of bools
package bitarr

const bitsInByte = 8

// Array represents an array of bits
type Array struct {
	b         []byte
	bitmapLen uint32
}

// NewArray returns a new initialised bitarray with a given size
func NewArray(bitmapLen uint32) *Array {
	return new(Array).Init(bitmapLen)
}

// Init initialises a bitarray with a given size, i.e. allocates the underlying byte array
func (s *Array) Init(bitmapLen uint32) *Array {
	s.b = make([]byte, 1+((bitmapLen-1)/bitsInByte)) // ceil(bitmapLen/bitsInByte)
	s.bitmapLen = bitmapLen
	return s
}

// Set sets a bit at a given position, i.e. a val with 0x1 will set it, 0x0 will clear it
// panics if pos is out-of-bounds
func (s *Array) Set(pos uint32, val byte) {
	if pos >= s.bitmapLen {
		panic("out of bounds :(")
	}
	whichByte := pos / bitsInByte
	whichPos := pos % bitsInByte
	val = val & 0x1
	n := byte(whichPos)
	mask := ^byte(0x1 << n)
	nByte := s.b[whichByte] & mask
	nByte |= byte(val << n)
	s.b[whichByte] = nByte
}

// Get gets a bit at a given position, returns
// panics if pos is out-of-bounds
func (s *Array) Get(pos uint32) byte {
	if pos >= s.bitmapLen {
		panic("out of bounds :(")
	}
	whichByte := pos / bitsInByte
	whichPos := pos % bitsInByte
	n := byte(whichPos)
	mask := byte(0x1 << n)
	bit := ((s.b[whichByte] & mask) >> n)
	return bit
}

// Len returns the length of the underlying bitarray (in bits)
func (s *Array) Len() uint32 {
	return s.bitmapLen
}
