package v1

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type MutexCounter struct {
	mutex sync.RWMutex
	v     int64
}

func (c *MutexCounter) Add(delta int64) {
	c.mutex.Lock()
	c.v += delta
	c.mutex.Unlock()
}

func (c *MutexCounter) Get() int64 {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.v
}

// Для запуска поменять название пакета на main
func main() {
	var (
		wg      sync.WaitGroup
		counter MutexCounter
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
