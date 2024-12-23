package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// В этом решении останавливаем горутину с помощью контекста
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine stopped!")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("some activity...")
			}
		}
	}(ctx)

	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second)
}
