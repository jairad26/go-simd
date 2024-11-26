package main

import "fmt"

// Declare the functions implemented in arm_int8.s
func addInt8Vec(result, a, b *int8, len int)
func subInt8Vec(result, a, b *int8, len int)
func dotInt8Vec(a, b *int8, len int) int32

func AddInt8Slices(a, b []int8) []int8 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]int8, len(a))

	addInt8Vec(&result[0], &a[0], &b[0], len(a))
	return result
}

func SubInt8Slices(a, b []int8) []int8 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	result := make([]int8, len(a))

	subInt8Vec(&result[0], &a[0], &b[0], len(a))
	return result
}

func DotInt8Slices(a, b []int8) int32 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slices must be same length: %d != %d", len(a), len(b)))
	}

	return dotInt8Vec(&a[0], &b[0], len(a))
}
