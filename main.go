package main

import (
	"fmt"
	"go-simd/simd_int8"
	"go-simd/simd_uint8"
)

func main() {
	uint8_a := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	uint8_b := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

	uintDot, _ := simd_uint8.DotVec(uint8_a, uint8_b)
	fmt.Println("Uint8 SIMD:", uintDot)

	int8_a := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	int8_b := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	intDot, _ := simd_int8.DotVec(int8_a, int8_b)
	fmt.Println("Int8 SIMD:", intDot)
}
