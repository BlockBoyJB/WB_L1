package main

import (
	"fmt"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	buckets := make(map[int][]float64)

	for _, t := range temperatures {
		key := int(t) - (int(t) % 10) // ключ - целое число, кратное 10
		buckets[key] = append(buckets[key], t)
	}
	fmt.Println(buckets)
}
