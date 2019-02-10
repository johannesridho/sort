package main

func Bubblesort(elements *[]int) {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(*elements)-1; i++ {
			if (*elements)[i] > (*elements)[i+1] {
				(*elements)[i], (*elements)[i+1] = (*elements)[i+1], (*elements)[i]
				swap = true
			}
		}
	}
}
