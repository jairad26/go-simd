package simd_int8

import "fmt"

// Declare the functions implemented in arm_int8.s
func AddVecSIMD(result, a, b *int8, len int)
func SubVecSIMD(result, a, b *int8, len int)
func DotVecSIMD16(a, b *int8, len int) int32
func DotVecSIMD32(a, b *int8, len int) int32
func DotVecSIMD64(a, b *int8, len int) int32

func AddVec(a, b []int8) ([]int8, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices must be same length: %d != %d", len(a), len(b))
	}

	if len(a) == 0 {
		return nil, fmt.Errorf("slices must have length greater than 0")
	}

	result := make([]int8, len(a))

	AddVecSIMD(&result[0], &a[0], &b[0], len(a))
	return result, nil
}

func SubVec(a, b []int8) ([]int8, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices must be same length: %d != %d", len(a), len(b))
	}

	if len(a) == 0 {
		return nil, fmt.Errorf("slices must have length greater than 0")
	}

	result := make([]int8, len(a))

	SubVecSIMD(&result[0], &a[0], &b[0], len(a))
	return result, nil
}

func DotVec(a, b []int8) (int32, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("slices must be same length: %d != %d", len(a), len(b))
	}

	if len(a) == 0 {
		return 0, fmt.Errorf("slices must have length greater than 0")
	}

	if len(a) < 32 {
		return DotVecSIMD16(&a[0], &b[0], len(a)), nil
	} else if len(a) < 64 {
		return DotVecSIMD32(&a[0], &b[0], len(a)), nil
	} else {
		return DotVecSIMD64(&a[0], &b[0], len(a)), nil
	}
}

func MultMatrix(a, b [][]int8) ([][]int32, error) {
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

			result[i][j], err = DotVec(a[i], column)
			if err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}
