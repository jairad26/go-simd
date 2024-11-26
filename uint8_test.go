package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// DotUInt8NEON is implemented in assembly

// Helper function to create test data
func generateUint8TestData(size int) ([]uint8, []uint8) {
	a := make([]uint8, size)
	b := make([]uint8, size)
	for i := range a {
		a[i] = uint8(rand.Intn(256))
		b[i] = uint8(rand.Intn(256))
	}
	return a, b
}

func BenchmarkUint8Add(b *testing.B) {
	sizes := []int{16, 100, 1000, 4096, 10000, 100000}

	for _, size := range sizes {
		a, v := generateUint8TestData(size)

		b.Run("Scalar-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				addUint8SlicesScalar(a, v)
			}
		})

		b.Run("SIMD-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				AddUint8Slices(a, v)
			}
		})
	}
}

func BenchmarkUint8Sub(b *testing.B) {
	sizes := []int{16, 100, 1000, 4096, 10000, 100000}

	for _, size := range sizes {
		a, v := generateUint8TestData(size)

		b.Run("Scalar-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				subUint8SlicesScalar(a, v)
			}
		})

		b.Run("SIMD-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				SubUint8Slices(a, v)
			}
		})
	}
}

func BenchmarkUint8DotProduct(b *testing.B) {
	// Test different sizes including non-multiple of 16
	sizes := []int{16, 100, 1000, 4096, 10000, 100000}

	for _, size := range sizes {
		// Generate test data outside the benchmark timing
		a, v := generateUint8TestData(size)

		b.Run("Scalar-"+fmt.Sprint(size), func(b *testing.B) {
			// Reset timer after any setup
			b.ResetTimer()
			// Clear any memory statistics from setup
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				dotUInt8Scalar(a, v)
			}
		})

		b.Run("SIMD-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				DotUInt8Slices(a, v)
			}
		})
	}
}

// TestCorrectness verifies that both implementations return the same results
func TestUint8Correctness(t *testing.T) {
	sizes := []int{15, 16, 17, 100, 1000}
	for _, size := range sizes {
		uint8_a, uint8_b := generateUint8TestData(size)

		uint8ScalarDot := dotUInt8Scalar(uint8_a, uint8_b)
		uint8SimdDot := DotUInt8Slices(uint8_a, uint8_b)

		if uint8ScalarDot != uint8SimdDot {
			t.Errorf("Size %d: Scalar: %d, SIMD: %d", size, uint8ScalarDot, uint8SimdDot)
		}

		uint8ScalarSum := addUint8SlicesScalar(uint8_a, uint8_b)
		uint8SimdSum := AddUint8Slices(uint8_a, uint8_b)

		for i := range uint8ScalarSum {
			if uint8ScalarSum[i] != uint8SimdSum[i] {
				t.Errorf("FOR ADD-> Size %d: Results don't match. Scalar: %d, SIMD: %d",
					size, uint8ScalarSum[i], uint8SimdSum[i])
			}
		}

		uint8ScalarDiff := subUint8SlicesScalar(uint8_a, uint8_b)
		uint8SimdDiff := SubUint8Slices(uint8_a, uint8_b)

		for i := range uint8ScalarDiff {
			if uint8ScalarDiff[i] != uint8SimdDiff[i] {
				t.Errorf("FOR SUB-> Size %d: Results don't match. Scalar: %d, SIMD: %d",
					size, uint8ScalarDiff[i], uint8SimdDiff[i])
			}
		}

	}
}
