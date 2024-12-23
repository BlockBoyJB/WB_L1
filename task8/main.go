package main

import (
	"fmt"
	"log"
)

func main() {
	var (
		number int64
		i      uint8
	)

	if _, err := fmt.Scan(&number, &i); err != nil {
		log.Fatal(err)
	}
	// xor операция. Как я понял, в задании нужно просто поменять бит на i-той позиции
	number = number ^ 1<<i
	fmt.Printf("number after xor: %d, binary %b", number, number)
}
