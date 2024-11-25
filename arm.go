package main

import "fmt"

// Declaration of our assembly function
func addUint8Vec(result, a, b *uint8, len int)
func addInt8Vec(result, a, b *int8, len int)
func dotInt8Vec(a, b *int8, len int) int32
func dotUint8Vec(a, b *uint8, len int) uint32

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

func AddInt8Slices(a, b []int8) []int8 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]int8, len(a))

	addInt8Vec(&result[0], &a[0], &b[0], len(a))
	return result
}

func DotUInt8Slices(a, b []uint8) uint32 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	return dotUint8Vec(&a[0], &b[0], len(a))
}

func DotInt8Slices(a, b []int8) int32 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	return dotInt8Vec(&a[0], &b[0], len(a))
}
