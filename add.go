package main

import "fmt"

// AddSlices adds two []uint8 slices using SIMD instructions
// Both slices must be the same length
func AddUint8Slices(a, b []uint8) []uint8 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]uint8, len(a))

	// Get pointer to start of each slice
	addUint8Vec(&result[0], &a[0], &b[0], len(a))
	return result
}

func AddUint8SlicesScalar(a, b []uint8) []uint8 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]uint8, len(a))

	for i := 0; i < len(a); i++ {
		result[i] = a[i] + b[i]
	}

	return result
}

func DotUInt8(a, b []uint8) uint32 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	return dotUint8Vec(&a[0], &b[0], len(a))
}

func DotUInt8Scalar(a, b []uint8) uint32 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	var result uint32
	for i := 0; i < len(a); i++ {
		result += uint32(a[i]) * uint32(b[i])
	}

	return result
}

// Declaration of our assembly function
func addUint8Vec(result, a, b *uint8, len int)
func dotUint8Vec(a, b *uint8, len int) uint32
