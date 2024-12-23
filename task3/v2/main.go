package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// решение с использованием atomic.Int64 для конкурентного вычисления
	a := []int64{2, 4, 6, 8, 10}

	var (
		count atomic.Int64
		wg    sync.WaitGroup
	)

	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go counter(&count, a[i], &wg)
	}
	wg.Wait()
	fmt.Println(count.Load())
}

func counter(count *atomic.Int64, value int64, wg *sync.WaitGroup) {
	count.Add(value * value)
	wg.Done()
}
