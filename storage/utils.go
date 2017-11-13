package storage

import (
	"fmt"
	"geodb/utils"
)

func extract8Bits(num uint8, offset uint8, length uint8) uint8 {

	retval, err := utils.ExtractBits(uint8(num), offset, length)

	if err != nil {
		fmt.Println(err, offset, length)
		panic("invalid bit extraction")
	}

	return retval.(uint8)

}

func extract32Bits(num int32, offset uint8, length uint8) uint32 {

	retval, err := utils.ExtractBits(uint32(num), offset, length)

	if err != nil {
		panic("invalid bit extraction")
	}

	return retval.(uint32)

}
