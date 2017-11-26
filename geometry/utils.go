package geometry

import (
	"math"
)

// extract8Bits extracts a set number of bits from an 8-bit number
// offset is determined from the left (ie leftmost bit is bit 0)
// offset 0 means start from bit 0
func extract8Bits(num int8, offset uint8, length uint8) uint8 {

	if length == 0 || length > 8 || offset > 7 {
		panic("invalid length or offset")
	}

	return uint8(num) & uint8(math.Pow(2, float64(8-offset))-1) >> (8 - offset - length)

}

// extract32Bits extracts a set number of bits from an 32-bit number
// offset is determined from the left (ie leftmost bit is bit 0)
func extract32Bits(num int32, offset uint8, length uint8) uint32 {

	if length == 0 || length > 32 || offset > 31 {
		panic("invalid length or offset")
	}

	return uint32(num) & uint32(math.Pow(2, float64(32-offset))-1) >> (32 - offset - length)

}
