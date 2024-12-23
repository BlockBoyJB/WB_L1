package main

import (
	counterV1 "WB_L1/task18/v1"
	counterV2 "WB_L1/task18/v2"
	counterV3 "WB_L1/task18/v3"
	"runtime"
	"sync"
	"testing"
)

func BenchmarkMutexCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := counterV1.MutexCounter{}

	for i := 0; i < runtime.NumCPU(); i++ {
		go func(ind int) {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				counter.Add(1)
			}
		}(i)
	}

	wg.Wait()
}

func BenchmarkAtomicCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := counterV2.AtomicCounter{}

	for i := 0; i < runtime.NumCPU(); i++ {
		go func(ind int) {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				counter.Add(1)
			}
		}(i)
	}

	wg.Wait()
}

func BenchmarkShardedCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	counter := counterV3.NewShardedAtomicCounter(runtime.NumCPU())

	for i := 0; i < runtime.NumCPU(); i++ {
		go func(ind int) {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				counter.Add(ind, 1)
			}
		}(i)
	}

	wg.Wait()
}
