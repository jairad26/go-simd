package simd_int8

import (
	"fmt"
	"math/rand"
	"testing"
)

func generateTestData(size int) ([]int8, []int8) {
	a := make([]int8, size)
	b := make([]int8, size)
	for i := range a {
		a[i] = int8(rand.Intn(256) - 128)
		b[i] = int8(rand.Intn(256) - 128)
	}
	return a, b
}

func generateMatrixTestData(rows, cols int) [][]int8 {
	a := make([][]int8, rows)
	for i := range a {
		a[i] = make([]int8, cols)
		for j := range a[i] {
			a[i][j] = int8(rand.Intn(256) - 128)
		}
	}
	return a
}

func BenchmarkAdd(b *testing.B) {
	sizes := []int{16, 100, 1000, 4096, 10000, 100000}

	for _, size := range sizes {
		a, v := generateTestData(size)

		b.Run("Scalar-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				_ = addSlicesScalar(a, v)
			}
		})

		b.Run("SIMD-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				_, _ = AddVec(a, v)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	sizes := []int{16, 100, 1000, 4096, 10000, 100000}

	for _, size := range sizes {
		a, v := generateTestData(size)

		b.Run("Scalar-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				_ = subSlicesScalar(a, v)
			}
		})

		b.Run("SIMD-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				_, _ = SubVec(a, v)
			}
		})
	}
}

func BenchmarkDotProduct(b *testing.B) {
	sizes := []int{16, 100, 1000, 4096, 10000, 100000}

	for _, size := range sizes {
		a, v := generateTestData(size)

		b.Run("Scalar-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				_ = dotScalar(a, v)
			}
		})

		b.Run("SIMD-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				_, _ = DotVec(a, v)
			}
		})
	}
}

func BenchmarkMatrixMult(b *testing.B) {
	sizes := []int{16, 100}

	for _, size := range sizes {
		a := generateMatrixTestData(size, size)
		v := generateMatrixTestData(size, size)

		b.Run("Scalar-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				_ = MultMatrixScalar(a, v)
			}
		})

		b.Run("SIMD-"+fmt.Sprint(size), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				_, _ = MultMatrix(a, v)
			}
		})
	}
}

func TestCorrectness(t *testing.T) {
	sizes := []int{15, 16, 17, 100, 496, 1000}
	for _, size := range sizes {

		int8_a, int8_b := generateTestData(size)

		int8ScalarDot := dotScalar(int8_a, int8_b)
		int8SimdDot, err := DotVec(int8_a, int8_b)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if int8ScalarDot != int8SimdDot {
			t.Errorf("Size %d: Scalar: %d, SIMD: %d", size, int8ScalarDot, int8SimdDot)
		}

		int8ScalarSum := addSlicesScalar(int8_a, int8_b)
		int8SimdSum, err := AddVec(int8_a, int8_b)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		for i := range int8ScalarSum {
			if int8ScalarSum[i] != int8SimdSum[i] {
				t.Errorf("FOR ADD-> Size %d: Results don't match. Scalar: %d, SIMD: %d",
					size, int8ScalarSum[i], int8SimdSum[i])
			}
		}

		int8ScalarDiff := subSlicesScalar(int8_a, int8_b)
		int8SimdDiff, err := SubVec(int8_a, int8_b)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		for i := range int8ScalarDiff {
			if int8ScalarDiff[i] != int8SimdDiff[i] {
				t.Errorf("FOR SUB-> Size %d: Results don't match. Scalar: %d, SIMD: %d",
					size, int8ScalarDiff[i], int8SimdDiff[i])
			}
		}

	}
}
