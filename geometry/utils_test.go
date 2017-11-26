package geometry

import "testing"

func TestExtract8Bits(t *testing.T) {

	tests8 := []struct {
		in     uint8
		offset uint8
		length uint8
		out    uint32
	}{
		{8, 2, 1, 0},
		{8, 0, 1, 0},
		{8, 0, 2, 0},
		{8, 0, 3, 0},
		{8, 0, 4, 0},
		{10, 2, 1, 0},
		{11, 2, 1, 0},
		{10, 0, 2, 0},
		{11, 0, 2, 0},
		{10, 2, 1, 0},
		{11, 2, 1, 0},
		{10, 1, 2, 0},
		{11, 1, 2, 0},
		{21, 2, 1, 0},
		{21, 2, 1, 0},
		{29, 2, 1, 0},
		{29, 1, 2, 0},
	}

	for _, tt := range tests8 {

		if bits := extract8Bits(int8(tt.in), tt.offset, tt.length); uint32(bits) != uint32(tt.out) {

			t.Errorf("in: %d, offset: %d, length: %d, exp: %d, got: %d", tt.in, tt.offset, tt.length, tt.out, bits)

		}
	}

}

func TestExtract8BitsZeroLength(t *testing.T) {

	defer func() {

		if r := recover(); r == nil {
			t.Error("expected a panic, but no panic was detected")
		}

	}()

	// Test failure cases
	extract8Bits(int8(10), 1, 0)

}

func TestExtract8BitsInvalidLength(t *testing.T) {

	defer func() {

		if r := recover(); r == nil {
			t.Error("expected a panic, but no panic was detected")
		}

	}()

	// Test failure cases
	extract8Bits(int8(10), 1, 9)

}

func TestExtract8BitsInvalidOffset(t *testing.T) {

	defer func() {

		if r := recover(); r == nil {
			t.Error("expected a panic, but no panic was detected")
		}

	}()

	// Test failure cases
	extract8Bits(int8(10), 8, 8)

}

func TestExtract32Bits(t *testing.T) {

	tests32 := []struct {
		in     uint32
		offset uint8
		length uint8
		out    uint32
	}{
		{100000000, 7, 5, 31},
		{1171645696, 12, 4, 5},
	}

	for _, tt := range tests32 {

		if bits := extract32Bits(int32(tt.in), tt.offset, tt.length); uint32(bits) != uint32(tt.out) {
			t.Errorf("in: %d, offset: %d, length: %d, exp: %d, got: %d", tt.in, tt.offset, tt.length, tt.out, bits)
		}
	}

}

func TestExtract32BitsZeroLength(t *testing.T) {

	defer func() {

		if r := recover(); r == nil {
			t.Error("expected a panic, but no panic was detected")
		}

	}()

	// Test failure cases
	extract32Bits(int32(10), 1, 0)

}

func TestExtract32BitsInvalidLength(t *testing.T) {

	defer func() {

		if r := recover(); r == nil {
			t.Error("expected a panic, but no panic was detected")
		}

	}()

	// Test failure cases
	extract32Bits(int32(10), 1, 33)

}

func TestExtract32BitsInvalidOffset(t *testing.T) {

	defer func() {

		if r := recover(); r == nil {
			t.Error("expected a panic, but no panic was detected")
		}

	}()

	// Test failure cases
	extract32Bits(int32(10), 32, 32)

}
