package simd_uint8

import "fmt"

// Declare the functions implemented in arm_uint8.s
//
//go:noescape
func AddVecSIMD(result, a, b *uint8, len int)

//go:noescape
func SubVecSIMD(result, a, b *uint8, len int)

//go:noescape
func DotVecSIMD16(a, b *uint8, len int) uint32

//go:noescape
func DotVecSIMD32(a, b *uint8, len int) uint32

//go:noescape
func DotVecSIMD64(a, b *uint8, len int) uint32

func AddVec(a, b []uint8) ([]uint8, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices must be same length: %d != %d", len(a), len(b))
	}

	if len(a) == 0 {
		return nil, fmt.Errorf("slices must have length greater than 0")
	}

	result := make([]uint8, len(a))

	// Get pointer to start of each slice
	AddVecSIMD(&result[0], &a[0], &b[0], len(a))
	return result, nil
}

func SubVec(a, b []uint8) ([]uint8, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices must be same length: %d != %d", len(a), len(b))
	}

	if len(a) == 0 {
		return nil, fmt.Errorf("slices must have length greater than 0")
	}

	result := make([]uint8, len(a))

	SubVecSIMD(&result[0], &a[0], &b[0], len(a))
	return result, nil
}

func DotVec(a, b []uint8) (uint32, error) {
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

func MultMatrix(a, b [][]uint8) ([][]uint32, error) {
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
			result[i][j], err = DotVec(a[i], column)
			if err != nil {
				return nil, err
			}
		}
	}
	return result, nil
}
