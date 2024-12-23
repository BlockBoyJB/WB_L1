package main

import (
	"fmt"
	"log"
)

func main() {
	var v any

	// Каким образом канал можно ввести с клавиатуры? =)
	if _, err := fmt.Scan(&v); err != nil {
		log.Fatal(err)
	}

	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan any:
		fmt.Println("chan any")
	}
}
