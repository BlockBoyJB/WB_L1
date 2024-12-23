package main

import (
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
	// В этом решении вместо привычного контекста используется обычный таймер
	t := time.After(time.Duration(n) * time.Second)

	go writer(t, c)
	go reader(t, c)

	<-t
	close(c)
}

func writer(t <-chan time.Time, c chan<- int) {
	for {
		select {
		case <-t:
			return
		default:
			n := rand.IntN(1000)
			c <- n
			time.Sleep(time.Second)
		}
	}
}

func reader(t <-chan time.Time, c <-chan int) {
	for {
		select {
		case <-t:
			return
		case v := <-c:
			fmt.Println(v)
		}
	}
}
