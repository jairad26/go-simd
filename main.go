package main

import (
	"fmt"
)

func main() {
	uint8_a := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	uint8_b := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	res := AddUint8Slices(uint8_a, uint8_b)
	fmt.Println(res)

	uintDot := DotUInt8Slices(uint8_a, uint8_b)
	fmt.Println("Uint8 SIMD:", uintDot)

	normalUintDot := dotUInt8Scalar(uint8_a, uint8_b)
	fmt.Println("Normal: ", normalUintDot)

	int8_a := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	int8_b := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	intDot := DotInt8Vec(int8_a, int8_b)
	fmt.Println("Int8 SIMD:", intDot)

	normalIntDot := dotInt8Scalar(int8_a, int8_b)
	fmt.Println("Normal: ", normalIntDot)
}
