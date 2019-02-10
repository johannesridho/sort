package main

import (
	"testing"
)

func BenchmarkBubblesort(b *testing.B) {
	elements := []int{3, 4, 2, 1, 4, 5, 1, 9, 2}
	for i := 0; i < b.N; i++ {
		Bubblesort(&elements)
	}
}
