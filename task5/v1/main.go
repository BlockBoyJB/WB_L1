package main

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"time"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatal(err)
	}

	c := make(chan int)
	// В этом решении решил использовать контекст с таймаутом, чтобы по истечении времени завершить выполнение всех горутин с этим контекстом
	// Можно было использовать context.WithCancel и просто time.Sleep, но меня осудят за такое
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(n))
	defer cancel()

	go writer(ctx, c)
	go reader(ctx, c)

	<-ctx.Done()
	close(c)
}

func writer(ctx context.Context, c chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			n := rand.IntN(1000)
			c <- n
			time.Sleep(time.Second)
		}
	}
}

func reader(ctx context.Context, c <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case v := <-c:
			fmt.Println(v)
		}
	}
}
