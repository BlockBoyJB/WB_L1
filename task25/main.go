package main

import "C"
import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	var d int
	if _, err := fmt.Scan(&d); err != nil {
		log.Fatal(err)
	}
	sleep1(d)
	sleep2(d)
	// можно было еще вызывать C функцию как 3 вариант
}

func sleep1(d int) {
	fmt.Println("sleep1 start")
	<-time.After(time.Duration(d) * time.Second)
	fmt.Println("sleep1 end")
}

func sleep2(d int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d)*time.Second)
	defer cancel()
	fmt.Println("sleep2 start")
	<-ctx.Done()
	fmt.Println("sleep2 end")
}
