package main

import "fmt"

// Declare the functions implemented in arm_int8.s
func addInt8Vec(result, a, b *int8, len int)
func subInt8Vec(result, a, b *int8, len int)
func dotInt8VecSIMD16(a, b *int8, len int) int32
func dotInt8VecSIMD32(a, b *int8, len int) int32
func dotInt8VecSIMD64(a, b *int8, len int) int32

func AddInt8Vec(a, b []int8) []int8 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]int8, len(a))

	addInt8Vec(&result[0], &a[0], &b[0], len(a))
	return result
}

func SubInt8Vec(a, b []int8) []int8 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]int8, len(a))

	subInt8Vec(&result[0], &a[0], &b[0], len(a))
	return result
}

func DotInt8Vec(a, b []int8) int32 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	if len(a) < 32 {
		return dotInt8VecSIMD16(&a[0], &b[0], len(a))
	} else if len(a) < 64 {
		return dotInt8VecSIMD32(&a[0], &b[0], len(a))
	} else {
		return dotInt8VecSIMD64(&a[0], &b[0], len(a))
	}
}

func MultInt8Matrix(a, b [][]int8) [][]int32 {
	if len(a[0]) != len(b) {
		panic(fmt.Sprintf("matrix a columns must be equal to matrix b rows: %d != %d", len(a[0]), len(b)))
	}

	result := make([][]int32, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = make([]int32, len(b[0]))
		column := make([]int8, len(b))
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				column[k] = b[k][j]
			}
			result[i][j] = DotInt8Vec(a[i], column)
		}
	}

	return result
}
