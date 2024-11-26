package main

import "fmt"

// Declare the functions implemented in arm_uint8.s
func addUint8Vec(result, a, b *uint8, len int)
func subUint8Vec(result, a, b *uint8, len int)
func dotUint8Vec(a, b *uint8, len int) uint32

func AddUint8Slices(a, b []uint8) []uint8 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]uint8, len(a))

	// Get pointer to start of each slice
	addUint8Vec(&result[0], &a[0], &b[0], len(a))
	return result
}

func SubUint8Slices(a, b []uint8) []uint8 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]uint8, len(a))

	subUint8Vec(&result[0], &a[0], &b[0], len(a))
	return result
}

func DotUInt8Slices(a, b []uint8) uint32 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	return dotUint8Vec(&a[0], &b[0], len(a))
}
