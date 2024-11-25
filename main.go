package main

import (
	"fmt"
)

func main() {
	uint8_a := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	uint8_b := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	res := AddUint8Slices(uint8_a, uint8_b)
	fmt.Println(res)

	uintDot := DotUInt8(uint8_a, uint8_b)
	fmt.Println(uintDot)

	// normalUintDot := DotUInt8Scalar(uint8_a, uint8_b)
	// fmt.Println(normalUintDot)
}
