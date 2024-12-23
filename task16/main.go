package main

import "fmt"

func partition(a []int, low, high int) int {
	pivot := a[high]

	i := low - 1

	for j := low; j < high; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}

func quickSort(a []int, low, high int) {
	if low < high {
		pi := partition(a, low, high)

		quickSort(a, low, pi-1)
		quickSort(a, pi+1, high)
	}
}

func main() {
	a := []int{10, 7, 8, 9, 1, 5}

	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
