package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	var (
		a, b, result big.Int
		t            string
	)
	fmt.Println("input 2 numbers with operand between: 2 + 1")
	if _, err := fmt.Scan(&a, &t, &b); err != nil {
		log.Fatal(err)
	}

	switch t {
	case "+":
		fmt.Println(result.Add(&a, &b))
	case "-":
		fmt.Println(result.Sub(&a, &b))
	case "*":
		fmt.Println(result.Mul(&a, &b))
	case "/":
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("division by zero!")
			}
		}()
		fmt.Println(result.Div(&a, &b))
	default:
		fmt.Println("unknown type of operation")
	}
}
