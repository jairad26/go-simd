package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func generateInt8TestData(size int) ([]int8, []int8) {
	a := make([]int8, size)
	b := make([]int8, size)
	for i := range a {
		a[i] = int8(rand.Intn(256) - 128)
		b[i] = int8(rand.Intn(256) - 128)
	}
	return a, b
}

func BenchmarkInt8Add(b *testing.B) {
	sizes := []int{16, 100, 1000, 4096, 10000, 100000}

	for _, size := range sizes {
		a, v := generateInt8TestData(size)

		b.Run("Scalar-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				addInt8SlicesScalar(a, v)
			}
		})

		b.Run("SIMD-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				AddInt8Slices(a, v)
			}
		})
	}
}

func BenchmarkInt8Sub(b *testing.B) {
	sizes := []int{16, 100, 1000, 4096, 10000, 100000}

	for _, size := range sizes {
		a, v := generateInt8TestData(size)

		b.Run("Scalar-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				subInt8SlicesScalar(a, v)
			}
		})

		b.Run("SIMD-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				SubInt8Slices(a, v)
			}
		})
	}
}

func BenchmarkInt8DotProduct(b *testing.B) {
	sizes := []int{16, 100, 1000, 4096, 10000, 100000}

	for _, size := range sizes {
		a, v := generateInt8TestData(size)

		b.Run("Scalar-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				dotInt8Scalar(a, v)
			}
		})

		b.Run("SIMD-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				DotInt8Slices(a, v)
			}
		})
	}
}

func TestInt8Correctness(t *testing.T) {
	sizes := []int{15, 16, 17, 100, 1000}
	for _, size := range sizes {

		int8_a, int8_b := generateInt8TestData(size)

		int8ScalarDot := dotInt8Scalar(int8_a, int8_b)
		int8SimdDot := DotInt8Slices(int8_a, int8_b)

		if int8ScalarDot != int8SimdDot {
			t.Errorf("Size %d: Scalar: %d, SIMD: %d", size, int8ScalarDot, int8SimdDot)
		}

		int8ScalarSum := addInt8SlicesScalar(int8_a, int8_b)
		int8SimdSum := AddInt8Slices(int8_a, int8_b)

		for i := range int8ScalarSum {
			if int8ScalarSum[i] != int8SimdSum[i] {
				t.Errorf("FOR ADD-> Size %d: Results don't match. Scalar: %d, SIMD: %d",
					size, int8ScalarSum[i], int8SimdSum[i])
			}
		}

		int8ScalarDiff := subInt8SlicesScalar(int8_a, int8_b)
		int8SimdDiff := SubInt8Slices(int8_a, int8_b)

		for i := range int8ScalarDiff {
			if int8ScalarDiff[i] != int8SimdDiff[i] {
				t.Errorf("FOR SUB-> Size %d: Results don't match. Scalar: %d, SIMD: %d",
					size, int8ScalarDiff[i], int8SimdDiff[i])
			}
		}

	}
}
