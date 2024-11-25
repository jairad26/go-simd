package main

import "fmt"

func addUint8SlicesScalar(a, b []uint8) []uint8 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]uint8, len(a))

	for i := 0; i < len(a); i++ {
		result[i] = a[i] + b[i]
	}

	return result
}

func addInt8SlicesScalar(a, b []int8) []int8 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]int8, len(a))

	for i := 0; i < len(a); i++ {
		result[i] = a[i] + b[i]
	}

	return result
}

func dotUInt8Scalar(a, b []uint8) uint32 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	var result uint32
	for i := 0; i < len(a); i++ {
		result += uint32(a[i]) * uint32(b[i])
	}

	return result
}

func dotInt8Scalar(a, b []int8) int32 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	var result int32
	for i := 0; i < len(a); i++ {
		result += int32(a[i]) * int32(b[i])
	}

	return result
}
