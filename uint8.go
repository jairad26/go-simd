package main

import "fmt"

// Declare the functions implemented in arm_uint8.s
func addUint8Vec(result, a, b *uint8, len int)
func subUint8Vec(result, a, b *uint8, len int)
func dotUint8VecSIMD16(a, b *uint8, len int) uint32
func dotUint8VecSIMD32(a, b *uint8, len int) uint32
func dotUint8VecSIMD64(a, b *uint8, len int) uint32

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

	if len(a) < 32 {
		return dotUint8VecSIMD16(&a[0], &b[0], len(a))
	} else if len(a) < 64 {
		return dotUint8VecSIMD32(&a[0], &b[0], len(a))
	} else {
		return dotUint8VecSIMD64(&a[0], &b[0], len(a))
	}
}

func MultUint8Matrix(a, b [][]uint8) [][]uint32 {
	if len(a[0]) != len(b) {
		panic(fmt.Sprintf("matrix a columns must be equal to matrix b rows: %d != %d", len(a[0]), len(b)))
	}

	result := make([][]uint32, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = make([]uint32, len(b[0]))
		column := make([]uint8, len(b))
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				column[k] = b[k][j]
			}
			result[i][j] = DotUInt8Slices(a[i], column)
		}
	}
	return result
}
