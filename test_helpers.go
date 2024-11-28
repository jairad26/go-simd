package main

import "fmt"

func addUint8SlicesScalar(a, b []uint8) []uint8 {
	if len(a) != len(b) {
		panic(fmt.Errorf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]uint8, len(a))

	for i := 0; i < len(a); i++ {
		result[i] = a[i] + b[i]
	}

	return result
}

func subUint8SlicesScalar(a, b []uint8) []uint8 {
	if len(a) != len(b) {
		panic(fmt.Errorf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]uint8, len(a))

	for i := 0; i < len(a); i++ {
		result[i] = a[i] - b[i]
	}

	return result
}

func dotUInt8Scalar(a, b []uint8) uint32 {
	if len(a) != len(b) {
		panic(fmt.Errorf("slices must be same length: %d != %d", len(a), len(b)))
	}

	var result uint32
	for i := 0; i < len(a); i++ {
		result += uint32(a[i]) * uint32(b[i])
	}

	return result
}

func addInt8SlicesScalar(a, b []int8) []int8 {
	if len(a) != len(b) {
		panic(fmt.Errorf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]int8, len(a))

	for i := 0; i < len(a); i++ {
		result[i] = a[i] + b[i]
	}

	return result
}

func subInt8SlicesScalar(a, b []int8) []int8 {
	if len(a) != len(b) {
		panic(fmt.Errorf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]int8, len(a))

	for i := 0; i < len(a); i++ {
		result[i] = a[i] - b[i]
	}

	return result
}

func dotInt8Scalar(a, b []int8) int32 {
	if len(a) != len(b) {
		panic(fmt.Errorf("slices must be same length: %d != %d", len(a), len(b)))
	}

	var result int32
	for i := 0; i < len(a); i++ {
		result += int32(a[i]) * int32(b[i])
	}

	return result
}

func MultInt8MatrixScalar(a, b [][]int8) [][]int32 {
	if len(a[0]) != len(b) {
		panic(fmt.Errorf("matrix a columns must be same length as matrix b rows: %d != %d", len(a[0]), len(b)))
	}

	if len(a) == 0 || len(b) == 0 {
		panic(fmt.Errorf("matrix a and b must have at least one row"))
	}

	result := make([][]int32, len(a))
	for i := range len(a) { // a rows
		result[i] = make([]int32, len(b[0])) // b columns
		for j := range len(b[0]) {           // b columns
			for k := 0; k < len(a[0]); k++ { // a columns
				result[i][j] += int32(a[i][k]) * int32(b[k][j]) // dot product
			}
		}
	}

	return result
}

func MultUint8MatrixScalar(a, b [][]uint8) [][]uint32 {
	if len(a[0]) != len(b) {
		panic(fmt.Errorf("matrix a columns must be same length as matrix b rows: %d != %d", len(a[0]), len(b)))
	}

	if len(a) == 0 || len(b) == 0 {
		panic(fmt.Errorf("matrix a and b must have at least one row"))
	}

	result := make([][]uint32, len(a))
	for i := range len(a) { // a rows
		result[i] = make([]uint32, len(b[0])) // b columns
		for j := range len(b[0]) {            // b columns
			for k := 0; k < len(a[0]); k++ { // a columns
				result[i][j] += uint32(a[i][k]) * uint32(b[k][j]) // dot product
			}
		}
	}

	return result
}
