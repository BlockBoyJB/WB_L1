package main

import (
	"fmt"
	"sync"
)

func main() {
	// В качестве примера будет мапа с целыми числами размером 10
	m := make(map[int]int, 10)

	wg := sync.WaitGroup{}
	mux := &sync.Mutex{}

	// 1000 горутин, каждая из которых будет увеличивать значение по ключу i % 10 на +1
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(m map[int]int, key int, mx *sync.Mutex) {
			mx.Lock()
			m[key] += 1
			mx.Unlock()
			wg.Done()
		}(m, i%10, mux)
	}

	wg.Wait()
	fmt.Println(m)
}
