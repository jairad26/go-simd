package main

import "fmt"

// Declare the functions implemented in arm_uint8.s
func addUint8Vec(result, a, b *uint8, len int)
func subUint8Vec(result, a, b *uint8, len int)
func dotUint8VecSIMD16(a, b *uint8, len int) uint32
func dotUint8VecSIMD32(a, b *uint8, len int) uint32
func dotUint8VecSIMD64(a, b *uint8, len int) uint32

func AddUint8Slices(a, b []uint8) ([]uint8, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices must be same length: %d != %d", len(a), len(b))
	}

	if len(a) == 0 {
		return nil, fmt.Errorf("slices must have length greater than 0")
	}

	result := make([]uint8, len(a))

	// Get pointer to start of each slice
	addUint8Vec(&result[0], &a[0], &b[0], len(a))
	return result, nil
}

func SubUint8Slices(a, b []uint8) ([]uint8, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices must be same length: %d != %d", len(a), len(b))
	}

	if len(a) == 0 {
		return nil, fmt.Errorf("slices must have length greater than 0")
	}

	result := make([]uint8, len(a))

	subUint8Vec(&result[0], &a[0], &b[0], len(a))
	return result, nil
}

func DotUInt8Slices(a, b []uint8) (uint32, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("slices must be same length: %d != %d", len(a), len(b))
	}

	if len(a) == 0 {
		return 0, fmt.Errorf("slices must have length greater than 0")
	}

	if len(a) < 32 {
		return dotUint8VecSIMD16(&a[0], &b[0], len(a)), nil
	} else if len(a) < 64 {
		return dotUint8VecSIMD32(&a[0], &b[0], len(a)), nil
	} else {
		return dotUint8VecSIMD64(&a[0], &b[0], len(a)), nil
	}
}

func MultUint8Matrix(a, b [][]uint8) ([][]uint32, error) {
	if len(a[0]) != len(b) {
		return nil, fmt.Errorf("matrix a columns must be equal to matrix b rows: %d != %d", len(a[0]), len(b))
	}

	if len(a) == 0 || len(b) == 0 || len(a[0]) == 0 || len(b[0]) == 0 {
		return nil, fmt.Errorf("matrices must have length greater than 0")
	}

	var err error

	result := make([][]uint32, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = make([]uint32, len(b[0]))
		column := make([]uint8, len(b))
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				column[k] = b[k][j]
			}
			result[i][j], err = DotUInt8Slices(a[i], column)
			if err != nil {
				return nil, err
			}
		}
	}
	return result, nil
}
