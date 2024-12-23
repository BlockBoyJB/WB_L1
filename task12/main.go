package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{}, len(a)) // мапа - отличное решение, когда нам нужна уникальность значений (т.е множество)
	for _, v := range a {
		set[v] = struct{}{} // Если в исходном слайсе есть повторения - они просто перезапишутся
	}

	var result []string
	for key, _ := range set {
		result = append(result, key)
	}
	fmt.Println("{" + strings.Join(result, ", ") + "}") // решил красиво сделать =)
}
