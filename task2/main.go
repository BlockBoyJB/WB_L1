package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [...]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for i := 0; i < len(numbers); i++ {
		wg.Add(1)
		go func(value int) {
			fmt.Println(value * value)
			wg.Done()
		}(numbers[i])
	}

	wg.Wait()
}
