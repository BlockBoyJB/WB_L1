package main

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatal(err)
	}

	// Все воркеры останавливаются путем отмены контекста
	p := &pool{
		numWorkers: n,
		job:        make(chan int),
	}

	ctx, cancel := context.WithCancel(context.Background())

	p.Start(ctx)
	// делаем так, чтобы остановить программу можно было через ctrl + c
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT) // Слушаем звуки сверчков...

	<-interrupt
	cancel()
	p.Stop()

}

type pool struct {
	numWorkers int
	wg         sync.WaitGroup
	job        chan int
}

func (p *pool) Start(ctx context.Context) {
	p.wg.Add(p.numWorkers)
	for i := 0; i < p.numWorkers; i++ {
		go p.worker(i, ctx)
	}
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				p.wg.Done()
				return
			default:
				n := rand.IntN(1000)
				p.job <- n
			}
		}
	}(ctx)
}

func (p *pool) worker(number int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			p.wg.Done()
			return
		case v := <-p.job:
			time.Sleep(time.Second)
			fmt.Printf("worker %d received value: %d\n", number, v)
		}
	}
}

func (p *pool) Stop() {
	fmt.Println("stop")
	p.wg.Wait()
}
