package main

import "fmt"

func main() {
	elements := []int{3, 4, 2, 1, 4, 5, 1, 9, 2}
	fmt.Println(elements)
	SelectionSort(elements)
	fmt.Println(elements)
}

func BubbleSort(elements []int) {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				swap = true
			}
		}
	}
}

func SelectionSort(elements []int) {
	for i := 0; i < len(elements); i++ {
		min := i
		for j := i + 1; j < len(elements); j++ {
			if elements[min] > elements[j] {
				min = j
			}
		}
		elements[min], elements[i] = elements[i], elements[min]
	}
}
