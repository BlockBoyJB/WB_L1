package main

import (
	"fmt"
	"sync"
)

func main() {
	// Решение с использованием обычного счетчика
	a := []int{2, 4, 6, 8, 10}

	var (
		count int
		wg    sync.WaitGroup
		mx    sync.Mutex
	)

	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go counter(&count, a[i], &mx, &wg)
	}
	wg.Wait()
	fmt.Println(count)
}

func counter(count *int, value int, mx *sync.Mutex, wg *sync.WaitGroup) {
	mx.Lock()
	*count += value * value
	mx.Unlock()
	wg.Done()
}
