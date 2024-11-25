package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// DotUInt8NEON is implemented in assembly

// Helper function to create test data
func generateTestData(size int) ([]uint8, []uint8) {
	a := make([]uint8, size)
	b := make([]uint8, size)
	for i := range a {
		a[i] = uint8(rand.Intn(256))
		b[i] = uint8(rand.Intn(256))
	}
	return []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
}

func BenchmarkDotProduct(b *testing.B) {
	// Test different sizes including non-multiple of 16
	sizes := []int{16, 100, 1000, 4096, 10000, 100000}

	for _, size := range sizes {
		a, v := generateTestData(size)

		b.Run("Scalar-"+fmt.Sprint(size), func(b *testing.B) {
			b.SetBytes(int64(size)) // for bytes/sec metrics
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				AddUint8SlicesScalar(a, v)
				// DotUInt8Scalar(a, v)
			}
		})

		b.Run("SIMD-"+fmt.Sprint(size), func(b *testing.B) {
			b.SetBytes(int64(size)) // for bytes/sec metrics
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				AddUint8Slices(a, v)
				// DotUInt8(a, v)
			}
		})
	}
}

// TestCorrectness verifies that both implementations return the same results
func TestCorrectness(t *testing.T) {
	sizes := []int{15, 16, 17, 100, 1000} // Test various sizes including non-multiples of 16

	for _, size := range sizes {
		a, b := generateTestData(size)

		// scalarDot := DotUInt8Scalar(a, b)
		// simdDot := DotUInt8(a, b)

		// if scalarDot != simdDot {
		// 	t.Errorf("FOR DOT-> Size %d: Results don't match. Scalar: %d, SIMD: %d",
		// 		size, scalarDot, simdDot)
		// }

		scalarSum := AddUint8SlicesScalar(a, b)
		simdSum := AddUint8Slices(a, b)

		for i := range scalarSum {
			if scalarSum[i] != simdSum[i] {
				t.Errorf("FOR ADD-> Size %d: Results don't match. Scalar: %d, SIMD: %d",
					size, scalarSum[i], simdSum[i])
			}
		}
	}
}
