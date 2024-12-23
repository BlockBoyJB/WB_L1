package main

import (
	"fmt"
	"log"
)

func main() {
	var (
		i int
		a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	)

	if _, err := fmt.Scan(&i); err != nil {
		log.Fatal(err)
	}

	// В срез от начала до i (не включая) добавляем срез от i+1 до конца исходного
	a = append(a[:i], a[i+1:]...)
	fmt.Println(a)
}
