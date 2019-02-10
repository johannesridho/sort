package main

import "fmt"

func main() {
	elements := []int{3, 4, 2, 1, 4, 5, 1, 9, 2}
	fmt.Println(elements)
	Bubblesort(&elements)
	fmt.Println(elements)
}
