package storage

import (
	"errors"
	"geodb/structs"
	"math"
)

func encodeLat(lat int32, retval []byte) error {

	if lat < -90000000 || lat > 90000000 {
		return errors.New("invalid latitude")
	}

	if lat >= 0 {
		lat = lat | int32(math.Pow(2, 27)) // flip on 28th bit
	} else {
		lat *= -1
	}

	retval[0] = byte(extract32Bits(lat, 4, 8))
	retval[1] = byte(extract32Bits(lat, 12, 8))
	retval[2] = byte(extract32Bits(lat, 20, 8))
	retval[3] = byte(extract32Bits(lat, 28, 4) << 4)

	return nil

}

func encodeLng(lng int32, retval []byte) error {

	if lng < -90000000 || lng > 90000000 {
		return errors.New("invalid longitude")
	}

	if lng >= 0 {
		lng = lng | int32(math.Pow(2, 28)) // flip on 29th bit
	} else {
		lng *= -1
	}

	retval[3] = byte(retval[3] | byte(extract32Bits(lng, 3, 4))) // bits 29-32: longitude except bit 29
	retval[4] = byte(extract32Bits(lng, 7, 8))                   // bits 33-40: longitude
	retval[5] = byte(extract32Bits(lng, 15, 8))                  // bits 41-48: longitude
	retval[6] = byte(extract32Bits(lng, 23, 8))                  // bits 49-56: longitude
	retval[7] = byte(extract32Bits(lng, 31, 1) << 7)             // bit 57: longitude

	return nil

}

func encodeElv(elv int32, retval []byte) error {

	if elv < 0 || elv > 32768 {
		return errors.New("invalid elevation")
	}

	retval[7] = byte(retval[7] | byte(extract32Bits(elv, 17, 7))) // bits 58-64: elevation
	retval[8] = byte(extract32Bits(elv, 24, 8))                   // bits 65-72: elevation

	return nil

}

// Encodes a point struct into a 9 byte array using the following mechanism:
// Bit 1: whether latitude is positive (1) or negative (0)
// Bit 2-28: latitude
// Bit 29: whether longitude is positive (1) or negative (0)
// Bit 30-57: longitude from 0-180 degrees
// Bit 58-72: elevation

func encodePoint(p structs.Point) []byte {

	retval := make([]byte, 9, 9)
	encodeLat(p.Lat, retval)
	encodeLng(p.Lng, retval)
	encodeElv(p.Elv, retval)

	return retval

}

func decodePoint(data []byte) structs.Point {

	p := structs.Point{
		Lat: 0,
		Lng: 0,
		Elv: 0,
	}

	elv := int32(255 & data[8])
	p.Elv = int32(extract8Bits(data[7], 1, 7))<<8 | elv

	// Reconstruct longitude
	p.Lng = int32(
		uint32(data[3])<<27 | // bits 27-25
			uint32(uint16(data[4]<<8)|uint16(data[5]))<<8 | // bits 24-9
			uint32((uint16(data[6])<<7)|uint16(extract8Bits(data[7], 0, 1)))) // bits 8-1

	if extract8Bits(data[3], 4, 1) == 0 {
		p.Lng *= -1
	}

	// Reconstruct latitude
	p.Lat = int32(data[0])<<20 | int32(data[1])<<12 | int32(data[2])<<4 | int32(data[3]>>4)

	if data[0]>>7 == 0 {
		p.Lat *= -1
	}

	return p

}
