package storage

import (
	"bytes"
	"fmt"
	"geodb/structs"
	"testing"
)

func TestEncodeLat(t *testing.T) {

	tests := []struct {
		in  int32
		exp []byte
	}{
		{0, []byte{128, 0, 0, 0, 0, 0, 0, 0, 0}},
		{1, []byte{128, 0, 0, 16, 0, 0, 0, 0, 0}},
		{-1, []byte{0, 0, 0, 16, 0, 0, 0, 0, 0}},
		{1000000, []byte{128, 244, 36, 0, 0, 0, 0, 0, 0}},
		{-1000000, []byte{0, 244, 36, 0, 0, 0, 0, 0, 0}},
		{45000000, []byte{170, 234, 84, 0, 0, 0, 0, 0, 0}},
		{-45000000, []byte{42, 234, 84, 0, 0, 0, 0, 0, 0}},
		{45000012, []byte{170, 234, 84, 192, 0, 0, 0, 0, 0}},
		{-45000012, []byte{42, 234, 84, 192, 0, 0, 0, 0, 0}},
		{89000006, []byte{212, 224, 132, 96, 0, 0, 0, 0, 0}},
		{-89000006, []byte{84, 224, 132, 96, 0, 0, 0, 0, 0}},
	}

	var bitstr []byte

	for _, tt := range tests {

		bitstr = make([]byte, 9)

		if encodeLat(tt.in, bitstr); bytes.Compare(tt.exp, bitstr) != 0 {
			t.Errorf("in: %d, exp: %b, out: %b", tt.in, tt.exp, bitstr)
		}
	}

	// Test boundary cases
	bitstr = make([]byte, 9)

	if err := encodeLat(90000001, bitstr); err == nil {
		t.Errorf("expected error but received none")
	}

	if err := encodeLat(-90000001, bitstr); err == nil {
		t.Errorf("expected error but received none")
	}

}

func TestEncodeLng(t *testing.T) {

	tests := []struct {
		in  int32
		exp []byte
	}{
		{0, []byte{212, 224, 132, 104, 0, 0, 0, 0, 0}},
		{1, []byte{212, 224, 132, 104, 0, 0, 0, 128, 0}},
		{-1, []byte{212, 224, 132, 96, 0, 0, 0, 128, 0}},
		{10000000, []byte{212, 224, 132, 104, 76, 75, 64, 0, 0}},
		{-10000000, []byte{212, 224, 132, 96, 76, 75, 64, 0, 0}},
		{89000006, []byte{212, 224, 132, 106, 167, 4, 35, 0, 0}},
		{-89000006, []byte{212, 224, 132, 98, 167, 4, 35, 0, 0}},
	}

	bitstr := make([]byte, 9)

	for _, tt := range tests {

		// Assume the input is 89000006
		bitstr = []byte{212, 224, 132, 96, 0, 0, 0, 0, 0}

		if encodeLng(tt.in, bitstr); bytes.Compare(tt.exp, bitstr) != 0 {
			t.Errorf("in: %d, exp: %d, out: %d", tt.in, tt.exp, bitstr)
		}
	}

	// Test boundary cases
	bitstr = make([]byte, 9)

	if err := encodeLng(180000001, bitstr); err == nil {
		t.Errorf("expected error but received none")
	}

	if err := encodeLng(-180000001, bitstr); err == nil {
		t.Errorf("expected error but received none")
	}

}

func TestEncodeElv(t *testing.T) {

	tests := []struct {
		in  int32
		exp []byte
	}{
		{0, []byte{212, 224, 132, 106, 167, 4, 35, 0, 0}},
		{1, []byte{212, 224, 132, 106, 167, 4, 35, 0, 1}},
		{1000, []byte{212, 224, 132, 106, 167, 4, 35, 3, 232}},
	}

	bitstr := make([]byte, 9)

	for _, tt := range tests {

		// Assume the input is 89000006
		bitstr = []byte{212, 224, 132, 106, 167, 4, 35, 0, 0}

		if encodeElv(tt.in, bitstr); bytes.Compare(tt.exp, bitstr) != 0 {
			t.Errorf("in: %d, exp: %d, out: %d", tt.in, tt.exp, bitstr)
		}
	}

	// Test boundary cases
	bitstr = make([]byte, 9)

	if err := encodeElv(32769, bitstr); err == nil {
		t.Errorf("expected error but received none")
	}

	if err := encodeElv(-1, bitstr); err == nil {
		t.Errorf("expected error but received none")
	}

}

func TestEncodePoint(t *testing.T) {

	t.Skip()

	input := structs.Point{
		Lat: 0,
		Lng: 0,
		Elv: 0,
	}

	encodedPoint := encodePoint(input)
	fmt.Printf("%b", encodedPoint)

}

func TestEncodeDecodePoint(t *testing.T) {

	t.Skip()
	input := structs.Point{
		Lat: 10,
		Lng: 10,
		Elv: 10,
	}

	encodedPoint := encodePoint(input)
	output := decodePoint(encodedPoint)

	if input.Lat != output.Lat {
		t.Error("input lat is not equal to output lat")
	} else if input.Lng != output.Lng {
		t.Error("input lng is not equal to output lng")
	} else if input.Elv != output.Elv {
		t.Error("input elv is not equal to output elv")
	}

}
