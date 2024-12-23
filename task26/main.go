package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var s string

	if _, err := fmt.Scan(&s); err != nil {
		log.Fatal(err)
	}

	fmt.Println(validate(s))
}

func validate(s string) bool {
	runes := []rune(strings.ToLower(s)) // Переводим в слайс рун, чтобы функция работала корректно
	symbols := make(map[rune]bool, len(runes))

	for i := 0; i < len(runes); i++ {
		v := runes[i]
		if symbols[v] {
			return false
		}
		symbols[v] = true
	}
	return true
}
