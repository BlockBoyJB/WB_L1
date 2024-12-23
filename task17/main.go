package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(binarySearch(a, 0, len(a)-1, 3))
}

func binarySearch(a []int, low, high, value int) int {
	for low <= high {
		mid := low + (high-low)/2

		if a[mid] == value {
			return mid
		}

		if a[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
