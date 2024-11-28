package main

import "fmt"

// Declare the functions implemented in arm_int8.s
func AddInt8VecSIMD(result, a, b *int8, len int)
func SubInt8VecSIMD(result, a, b *int8, len int)
func DotInt8VecSIMD16(a, b *int8, len int) int32
func DotInt8VecSIMD32(a, b *int8, len int) int32
func DotInt8VecSIMD64(a, b *int8, len int) int32

func AddInt8Vec(a, b []int8) ([]int8, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices must be same length: %d != %d", len(a), len(b))
	}

	if len(a) == 0 {
		return nil, fmt.Errorf("slices must have length greater than 0")
	}

	result := make([]int8, len(a))

	AddInt8VecSIMD(&result[0], &a[0], &b[0], len(a))
	return result, nil
}

func SubInt8Vec(a, b []int8) ([]int8, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices must be same length: %d != %d", len(a), len(b))
	}

	if len(a) == 0 {
		return nil, fmt.Errorf("slices must have length greater than 0")
	}

	result := make([]int8, len(a))

	SubInt8VecSIMD(&result[0], &a[0], &b[0], len(a))
	return result, nil
}

func DotInt8Vec(a, b []int8) (int32, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("slices must be same length: %d != %d", len(a), len(b))
	}

	if len(a) == 0 {
		return 0, fmt.Errorf("slices must have length greater than 0")
	}

	if len(a) < 32 {
		return DotInt8VecSIMD16(&a[0], &b[0], len(a)), nil
	} else if len(a) < 64 {
		return DotInt8VecSIMD32(&a[0], &b[0], len(a)), nil
	} else {
		return DotInt8VecSIMD64(&a[0], &b[0], len(a)), nil
	}
}

func MultInt8Matrix(a, b [][]int8) ([][]int32, error) {
	if len(a[0]) != len(b) {
		return nil, fmt.Errorf("matrix a columns must be equal to matrix b rows: %d != %d", len(a[0]), len(b))
	}
	var err error

	result := make([][]int32, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = make([]int32, len(b[0]))
		column := make([]int8, len(b))
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				column[k] = b[k][j]
			}

			result[i][j], err = DotInt8Vec(a[i], column)
			if err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}
