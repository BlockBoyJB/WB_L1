package v2

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type AtomicCounter struct {
	v atomic.Int64
}

func (c *AtomicCounter) Add(delta int64) {
	c.v.Add(delta)
}

func (c *AtomicCounter) Get() int64 {
	return c.v.Load()
}

// Для запуска поменять название пакета на main
func main() {
	var (
		wg      sync.WaitGroup
		counter AtomicCounter
	)
	start := time.Now()
	n := runtime.NumCPU()

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < 1_000_000; j++ {
				counter.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter.Get() / int64(n))
	fmt.Println(time.Since(start))
}
