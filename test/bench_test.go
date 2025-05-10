package main

import (
	"fmt"
	"testing"
)

// Without capacity
func BenchmarkMapWithoutCapacity(b *testing.B) {
	fmt.Println("")
	fmt.Println("BenchmarkMapWithoutCapacity b.N:", b.N)
	fmt.Println("")

	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		for j := 0; j < 10000; j++ {
			m[j] = j
		}
	}
}

// With capacity
func BenchmarkMapWithCapacity(b *testing.B) {
	fmt.Println("")
	fmt.Println("BenchmarkMapWithCapacity b.N:", b.N)
	fmt.Println("")

	for i := 0; i < b.N; i++ {
		m := make(map[int]int, 10000)
		for j := 0; j < 10000; j++ {
			m[j] = j
		}
	}
}
