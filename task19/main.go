package main

import (
	"fmt"
	"log"
)

func main() {
	//fmt.Println(reverse("foobar"))
	//
	//fmt.Println(reverse("главрыба"))
	//
	//fmt.Println(reverse("привет")) // вывод - кодовая фраза из смешариков =)

	var s string
	if _, err := fmt.Scan(&s); err != nil {
		log.Fatal(err)
	}

	fmt.Println(reverse(s))
}

func reverse(s string) string {
	r := []rune(s) // В задании сказано, что символы unicode - значит надо превратить их в руны =)
	result := make([]rune, len(r))

	for i := len(r) - 1; i >= 0; i-- {
		result = append(result, r[i])
	}
	return string(result)
}
