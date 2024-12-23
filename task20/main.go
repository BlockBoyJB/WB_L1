package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// простой fmt.Scan не подходит, потому что пробелы все сломали =(
	line, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	s := string(line)

	fmt.Println(reverse(s))
}

func reverse(s string) (result string) {
	sp := strings.Split(s, " ")

	for i := len(sp) - 1; i > 0; i-- {
		result += sp[i] + " "
	}
	result += sp[0]
	return
}
