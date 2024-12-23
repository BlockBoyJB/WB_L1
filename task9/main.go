package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	job := make(chan int)
	result := make(chan int)

	go producer(ctx, job)
	go squaring(ctx, job, result)
	go consumer(ctx, result)

	time.Sleep(time.Second * 5)
	cancel()
	time.Sleep(time.Second)
}

func producer(ctx context.Context, job chan<- int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("producer stopped")
			return
		default:
			n := rand.IntN(1000)
			job <- n
			time.Sleep(time.Second)
		}
	}
}

func squaring(ctx context.Context, job <-chan int, result chan<- int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("squaring stopped")
			return
		case v := <-job:
			fmt.Println("squaring received value", v)
			result <- v * v
		}
	}
}

func consumer(ctx context.Context, result <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("consumer stopped")
			return
		case v := <-result:
			fmt.Println("consumer received value", v)
		}
	}
}
