package main

import "testing"

func BenchmarkBubbleSort(b *testing.B) {
	elements := []int{3, 4, 2, 1, 4, 5, 1, 9, 2}
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	elements := []int{3, 4, 2, 1, 4, 5, 1, 9, 2}
	for i := 0; i < b.N; i++ {
		SelectionSort(elements)
	}
}
