//go:build !race

package simd_int8

import (
	"fmt"
	"math/rand"
	"testing"
	"unsafe"
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

//go:nocheckptr
func TestMemoryAlignmentRigorous(t *testing.T) {
	testSizes := []struct {
		name       string
		size       int
		alignments []int
	}{
		{"SIMD16", 17, []int{1, 3, 7, 11, 13, 15}},
		{"SIMD32", 33, []int{1, 5, 11, 17, 23, 29, 31}},
		{"SIMD64", 65, []int{1, 7, 13, 15, 23, 31, 37, 47, 59, 63}},
		{"SIMD128", 129, []int{1, 7, 13, 15, 23, 31, 37, 47, 59, 63}},
		{"VERY_LARGE", 1024, []int{1, 7, 13, 15, 23, 31, 37, 47, 59, 63}},
		{"VERY_VERY_LARGE", 10000, []int{1, 7, 13, 15, 23, 31, 37, 47, 59, 63}},
	}

	for _, ts := range testSizes {
		t.Run(ts.name, func(t *testing.T) {
			rawBuffer := make([]byte, 16384+64)
			baseAddr := uintptr(unsafe.Pointer(&rawBuffer[0]))
			padding := (64 - (baseAddr % 64)) % 64

			for _, misalign := range ts.alignments {
				misalignedAddr := baseAddr + padding + uintptr(misalign)
				misalignedPtr := unsafe.Pointer(misalignedAddr)

				a := unsafe.Slice((*byte)(misalignedPtr), ts.size)
				b := unsafe.Slice((*byte)(misalignedPtr), ts.size)

				for i := 0; i < ts.size; i++ {
					a[i] = byte(i % 128)
					b[i] = byte(i % 128)
				}

				aAddr := uintptr(unsafe.Pointer(&a[0]))
				bAddr := uintptr(unsafe.Pointer(&b[0]))

				if aAddr%16 == 0 || bAddr%16 == 0 {
					t.Errorf("%s still aligned at offset %d: a:%x (mod 16: %d) b:%x (mod 16: %d)",
						ts.name, misalign, aAddr, aAddr%16, bAddr, bAddr%16)
				}

				aInt8 := *(*[]int8)(unsafe.Pointer(&a))
				bInt8 := *(*[]int8)(unsafe.Pointer(&b))
				expected := dotScalar(aInt8, bInt8)
				result, err := DotVec(aInt8, bInt8)

				if err != nil {
					t.Fatalf("%s error at offset %d: %v", ts.name, misalign, err)
				}

				if result != expected {
					t.Errorf("%s offset %d: misaligned addresses failed: got %d, want %d\na:%x (mod 32: %d)\nb:%x (mod 32: %d)",
						ts.name, misalign, result, expected, aAddr, aAddr%32, bAddr, bAddr%32)
				}
			}
		})
	}
}
