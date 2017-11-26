package geometry

import (
	"errors"
	"geodb/storage"
	"math"
)

var storageWrite = storage.Write

// Write writes a point to the storage layer
func (p *Point) Write() error {

	pr, err := newRecord(*p)

	if err != nil {
		return err
	}

	if err := storageWrite(toBytes(pr)); err != nil {
		return err
	}

	return nil

}

// Encode encodes a point struct into a 9 byte array using the following mechanism:
// Bit 1: whether latitude is positive (1) or negative (0)
// Bit 2-28: latitude
// Bit 29: whether longitude is positive (1) or negative (0)
// Bit 30-57: longitude from 0-180 degrees
// Bit 58-72: elevation
func (p *Point) Encode() ([]byte, error) {

	var err error
	retval := make([]byte, 9)

	if retval, err = encodeLat(p.Lat, retval); err != nil {
		return nil, err
	}

	if retval, err = encodeLng(p.Lng, retval); err != nil {
		return nil, err
	}

	if retval, err = encodeElv(p.Elv, retval); err != nil {
		return nil, err
	}

	return retval, nil

}

// Basic rules when conversing with this code
// 1. Bits offsets are from right to left
// 2. Bits offsets start from 0 (eg bit 0 is rightmost bit)

// encodeLat encodes latitude as follows:
// Bit 1: whether latitude is positive (1) or negative (0)
// Bit 2-28: latitude
// Latitude is encoded using 28 bits
func encodeLat(lat int32, data []byte) ([]byte, error) {

	retval := make([]byte, len(data))
	copy(retval, data)

	// Check for valid latitude (positive/negative 90 degrees)
	if lat < -90000000 || lat > 90000000 {
		return retval, errors.New("invalid latitude")
	}

	if lat >= 0 {

		// Positive latitude have 28th bit set to 1
		lat = lat | int32(math.Pow(2, 27)) // flip on 28th bit

	} else {

		// Negative latitude have 28th bit set to 0
		lat *= -1
	}

	retval[0] = byte(extract32Bits(lat, 4, 8))       // bits 0-7
	retval[1] = byte(extract32Bits(lat, 12, 8))      // bits 8-15
	retval[2] = byte(extract32Bits(lat, 20, 8))      // bits 16-31
	retval[3] = byte(extract32Bits(lat, 28, 4) << 4) // bits 32-35

	return retval, nil

}

// encodeLat encodes a latitude to a byte array
// encodeLng encodes longitude as follows:
// Bit 29: whether longitude is positive (1) or negative (0)
// Bit 30-57: longitude from 0-180 degrees
// Latitude is encoded using 29 bits
func encodeLng(lng int32, data []byte) ([]byte, error) {

	retval := make([]byte, len(data))
	copy(retval, data)

	if lng < -180000000 || lng > 180000000 {
		return retval, errors.New("invalid longitude")
	}

	if lng >= 0 {

		// Positive longitude have 29th bit set to 1
		lng = lng | int32(math.Pow(2, 28)) // flip on 29th bit

	} else {

		// Negative longitude have 29th bit set to 0
		lng *= -1

	}

	retval[3] = retval[3] | byte(extract32Bits(lng, 3, 4)) // bits 29-32: longitude except bit 29
	retval[4] = byte(extract32Bits(lng, 7, 8))             // bits 33-40: longitude
	retval[5] = byte(extract32Bits(lng, 15, 8))            // bits 41-48: longitude
	retval[6] = byte(extract32Bits(lng, 23, 8))            // bits 49-56: longitude
	retval[7] = byte(extract32Bits(lng, 31, 1) << 7)       // bit 57: longitude

	return retval, nil

}

// encodeElv encodes elevation as follows:
// Bit 58-72: elevation
// Elevation is encoded using 15 bits and therefore can only encode between 0 and 32767 (inclusive)
func encodeElv(elv int32, data []byte) ([]byte, error) {

	retval := make([]byte, len(data))
	copy(retval, data)

	if elv < 0 || elv > 32767 {
		return retval, errors.New("invalid elevation")
	}

	retval[7] = retval[7] | byte(extract32Bits(elv, 17, 7)) // bits 58-64: elevation
	retval[8] = byte(extract32Bits(elv, 24, 8))             // bits 65-72: elevation

	return retval, nil

}

// decodeLat decodes a given byte string and returns the extracted latitude
func decodeLat(data []byte) (int32, error) {

	lat := int32(-1)

	if data[0]>>7 == 1 {

		lat = 1
		data[0] &= 127 // flip 8th bit back to 0

	}

	// Reconstruct latitude
	lat *= int32(data[0])<<20 | int32(data[1])<<12 | int32(data[2])<<4 | int32(data[3]>>4)

	if lat > 90000000 || lat < -90000000 {
		return lat, errors.New("invalid latitude ")
	}

	return lat, nil

}

// decodeLng decodes the bytestring and returns the extracted longitude
func decodeLng(data []byte) (int32, error) {

	lng := int32(-1)
	data[3] &= 15 // zero out top 4 bits as the top four are used in lat

	if data[3]>>3 == 1 {
		lng = 1
		data[3] &= 7 // flip 4th bit back to 0
	}

	// Reconstruct longitude
	lng *= int32(
		uint32(data[3])<<25 | // bit 27-25
			uint32(data[4])<<17 | // bit 24-17
			uint32(data[5])<<9 | // bit 16-9
			uint32(data[6])<<1 | // bit 8-1
			uint32(extract8Bits(int8(data[7]), 0, 1))) // bit 0

	if lng > 180000000 || lng < -180000000 {
		return lng, errors.New("invalid longitude")
	}

	return lng, nil

}

// decodeElv decodes the bytestring and returns the extracted elevation
func decodeElv(data []byte) int32 {

	elv := int32(data[8])                                    // bit 7-0
	return int32(extract8Bits(int8(data[7]), 1, 7))<<8 | elv // bit 14-8

}

// decodePoint decodes the bytestring and an entire point struct
func decodePoint(data []byte) (Point, error) {

	var p Point
	var err error

	p.Lat, err = decodeLat(data)

	if err != nil {
		return p, err
	}

	p.Lng, err = decodeLng(data)

	if err != nil {
		return p, err
	}

	p.Elv = decodeElv(data)

	return p, nil

}
