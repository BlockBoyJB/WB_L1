// Идею с шардированием подсмотрел у Владимира Балуна в ролике про оптимизацию кода, где как раз раскрывалось конкурентное увеличение счетчика
package v3

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type shard struct {
	v atomic.Int64
	_ [56]byte // Выравнивание структуры до длины кэш линии (8 байт для атомика + 56 байт выравнивания = 64 байта - одна кэш линия) (значение может отличаться в зависимости от архитектуры)
}

type ShardedAtomicCounter struct {
	shards []*shard
}

func NewShardedAtomicCounter(numShards int) *ShardedAtomicCounter {
	c := &ShardedAtomicCounter{
		shards: make([]*shard, numShards),
	}
	for i := 0; i < numShards; i++ {
		c.shards[i] = &shard{}
	}
	return c
}

func (c *ShardedAtomicCounter) Add(shard int, delta int64) {
	c.shards[shard].v.Add(delta)
}

func (c *ShardedAtomicCounter) Get() (v int64) {
	for _, s := range c.shards {
		v += s.v.Load()
	}
	return
}

// Для запуска поменять название пакета на main
func main() {
	wg := sync.WaitGroup{}
	start := time.Now()
	n := runtime.NumCPU()
	counter := NewShardedAtomicCounter(n)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(ind int) {
			for j := 0; j < 1_000_000; j++ {
				counter.Add(ind, 1)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(counter.Get() / int64(n))
	fmt.Println(time.Since(start))
}
