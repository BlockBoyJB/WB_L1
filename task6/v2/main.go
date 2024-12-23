package main

import (
	"fmt"
	"time"
)

func main() {
	// В этом решении останавливаем горутину с помощью закрытия канала
	ch := make(chan struct{})

	go func(c <-chan struct{}) {
		for {
			select {
			case <-c:
				// Когда мы закрываем канал, то чтение из него вернет zero value
				fmt.Println("goroutine stopped!")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("some activity...")
			}
		}
	}(ch)

	time.Sleep(time.Second * 3)
	close(ch)
	time.Sleep(time.Second)
}
