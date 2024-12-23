package main

import (
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

	// Все воркеры останавливаются путем закрытия канала "stop". Чтение из него вернет значение по умолчанию
	// В select case, где слушается канал, произойдет остановка, потому что операция чтения из канала перестанет блокироваться
	p := &pool{
		numWorkers: n,
		job:        make(chan int),
		stop:       make(chan struct{}),
	}

	p.Start()
	// делаем так, чтобы остановить программу можно было через ctrl + c
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT) // Слушаем звуки сверчков...

	<-interrupt
	p.Stop()
}

type pool struct {
	numWorkers int
	wg         sync.WaitGroup
	job        chan int // Основной канал
	stop       chan struct{}
}

func (p *pool) Start() {
	p.wg.Add(p.numWorkers + 1)
	for i := 0; i < p.numWorkers; i++ {
		go p.worker(i)
	}

	// Горутина, которая пушит значения в канал
	go func() {
		for {
			select {
			case <-p.stop:
				p.wg.Done()
				return
			default:
				n := rand.IntN(1000)
				p.job <- n
			}
		}
	}()
}

func (p *pool) worker(number int) {
	for {
		select {
		case <-p.stop:
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
	close(p.stop)
	p.wg.Wait()
}
