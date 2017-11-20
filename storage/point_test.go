package storage

import (
	"bytes"
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
		{90000000, []byte{213, 212, 168, 0, 0, 0, 0, 0, 0}}, // boundary case
		{-90000000, []byte{85, 212, 168, 0, 0, 0, 0, 0, 0}}, // boundary case
	}

	var bitstr []byte

	for _, tt := range tests {

		bitstr = make([]byte, 9)

		if bitstr, _ := encodeLat(tt.in, bitstr); bytes.Compare(tt.exp, bitstr) != 0 {
			t.Errorf("in: %d, exp: %d, out: %d", tt.in, tt.exp, bitstr)
		}
	}

	// Test boundary cases
	bitstr = make([]byte, 9)

	if _, err := encodeLat(90000001, bitstr); err == nil {
		t.Errorf("expected error but received none")
	}

	if _, err := encodeLat(-90000001, bitstr); err == nil {
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
		{180000000, []byte{212, 224, 132, 109, 93, 74, 128, 0, 0}},  // boundary case
		{-180000000, []byte{212, 224, 132, 101, 93, 74, 128, 0, 0}}, // boundary case
	}

	bitstr := make([]byte, 9)

	for _, tt := range tests {

		// Assume the input is 89000006
		bitstr = []byte{212, 224, 132, 96, 0, 0, 0, 0, 0}

		if bitstr, _ := encodeLng(tt.in, bitstr); bytes.Compare(tt.exp, bitstr) != 0 {
			t.Errorf("in: %d, exp: %d, out: %d", tt.in, tt.exp, bitstr)
		}
	}

	// Test boundary cases
	bitstr = []byte{212, 224, 132, 96, 0, 0, 0, 0, 0}

	if _, err := encodeLng(180000001, bitstr); err == nil {
		t.Errorf("expected error but received none")
	}

	if _, err := encodeLng(-180000001, bitstr); err == nil {
		t.Errorf("expected error but received none")
	}

}

func TestEncodeElv(t *testing.T) {

	tests := []struct {
		in  int32
		exp []byte
	}{
		{0, []byte{212, 224, 132, 106, 167, 4, 35, 0, 0}}, // boundary case
		{1, []byte{212, 224, 132, 106, 167, 4, 35, 0, 1}},
		{1000, []byte{212, 224, 132, 106, 167, 4, 35, 3, 232}},
		{32767, []byte{212, 224, 132, 106, 167, 4, 35, 127, 255}}, // boundary case
	}

	bitstr := make([]byte, 9)

	for _, tt := range tests {

		// Assume the input is 89000006
		bitstr = []byte{212, 224, 132, 106, 167, 4, 35, 0, 0}

		if bitstr, _ := encodeElv(tt.in, bitstr); bytes.Compare(tt.exp, bitstr) != 0 {
			t.Errorf("in: %d, exp: %d, out: %d", tt.in, tt.exp, bitstr)
		}
	}

	// Test boundary cases
	bitstr = []byte{212, 224, 132, 106, 167, 4, 35, 0, 0}

	if _, err := encodeElv(32769, bitstr); err == nil {
		t.Errorf("expected error but received none")
	}

	if _, err := encodeElv(-1, bitstr); err == nil {
		t.Errorf("expected error but received none")
	}

}

func TestEncodePoint(t *testing.T) {

	p := structs.Point{
		Lat: 10,
		Lng: 10,
		Elv: 10,
	}

	if _, err := encodePoint(p); err != nil {
		t.Error("encoded 0,0,0, received err")
	}

	p.Lat = 90000001

	if _, err := encodePoint(p); err == nil {
		t.Error("encoded invalid lat, received no err")
	}

	p.Lat = 0
	p.Lng = 180000001

	if _, err := encodePoint(p); err == nil {
		t.Error("encoded invalid lng, received no err")
	}

	p.Lng = 0
	p.Elv = 32769

	if _, err := encodePoint(p); err == nil {
		t.Error("encoded invalid elv, received no err")
	}

}

func TestDecodeLat(t *testing.T) {

	bytestr := []byte{212, 224, 132, 106, 167, 4, 35, 3, 232}
	lat, err := decodeLat(bytestr)

	if err != nil {
		t.Fatal("unable to decode valid latitude")
	}

	if lat != 89000006 {
		t.Errorf("expected 89000006, got %d", lat)
	}

	bytestr = []byte{85, 212, 168, 16, 0, 0, 0, 0}
	lat, err = decodeLat(bytestr)

	if err == nil {
		t.Error("able to decode invalid latitude")
	}

}

func TestDecodeLng(t *testing.T) {

	bytestr := []byte{212, 224, 132, 106, 167, 4, 35, 3, 232}
	lng, err := decodeLng(bytestr)

	if err != nil {
		t.Fatal("unable to encode decode valid longitude")
	}

	if lng != 89000006 {
		t.Errorf("expected 89000006, got %d", lng)
	}

	bytestr = []byte{212, 224, 132, 109, 93, 74, 128, 128, 0}
	lng, err = decodeLng(bytestr)

	if err == nil {
		t.Error("able to decode invalid longitude")
	}

}

func TestDecodeElv(t *testing.T) {

	bytestr := []byte{212, 224, 132, 106, 167, 4, 35, 3, 232}
	elv := decodeElv(bytestr)

	if elv != 1000 {
		t.Errorf("expected 1000, got %d", elv)
	}

}

func TestDecodeInvalidPoint(t *testing.T) {

	bytestr := []byte{85, 212, 168, 16, 0, 0, 0, 0}
	_, err := decodePoint(bytestr)

	if err == nil {
		t.Error("able to decode invalid bytestr")
	}

	bytestr = []byte{212, 224, 132, 109, 93, 74, 128, 128, 0}
	_, err = decodePoint(bytestr)

	if err == nil {
		t.Error("able to decode invalid bytestr")
	}

}

func TestEncodeDecodePoint(t *testing.T) {

	tests := []struct {
		lat int32
		lng int32
		elv int32
	}{
		{0, 0, 0},
		{0, 0, 1},
		{10000000, 10000000, 1},
		{-10000000, 10000000, 1},
		{10000000, 10000000, 10},
		{-10000000, 10000000, 10},
		{90000000, 180000000, 32767},   // boundary case
		{-90000000, -180000000, 32767}, // boundary case
	}

	p := structs.Point{
		Lat: 0,
		Lng: 0,
		Elv: 0,
	}

	for _, tt := range tests {

		p = structs.Point{
			Lat: tt.lat,
			Lng: tt.lng,
			Elv: tt.elv,
		}

		ep, err := encodePoint(p)

		if err != nil {
			t.Fatal("unable to encode valid point")
		}

		dp, err := decodePoint(ep)

		if err != nil {
			t.Fatal("unable to decode valid point")
		}

		if p.Lat != dp.Lat {
			t.Error("input lat is not equal to output lat")
		} else if p.Lng != dp.Lng {
			t.Error("input lng is not equal to output lng")
		} else if p.Elv != dp.Elv {
			t.Error("input elv is not equal to output elv")
		}
	}
}
