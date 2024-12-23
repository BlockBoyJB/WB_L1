package main

import (
	"fmt"
	"time"
)

func main() {
	// В этом решении горутина будет слушать канал с временем
	timer := time.After(time.Second * 3)

	go func(t <-chan time.Time) {
		for {
			select {
			case <-t:
				fmt.Println("goroutine stopped")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("some activity...")
			}
		}
	}(timer)

	time.Sleep(time.Second * 4)
}
