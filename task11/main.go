package main

import "fmt"

func main() {
	a1 := []int{1, 3, 5, 7, 9}
	a2 := []int{2, 4, 6, 8, 10}

	fmt.Println(intersection(a1, a2)) // []

	b1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13}
	b2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	fmt.Println(intersection(b1, b2)) // [1 2 3 4 5 6 7 8 9 10]

	c1 := []string{"go", "golang", "love_golang", "foobar", "abc", "cba", "foo", "bar"}
	c2 := []string{"golang", "love_golang", "gopher", "foobar", "bar", "abc", "hello", "world"}

	fmt.Println(intersection(c1, c2)) // [golang love_golang foobar bar abc]
}

// Пересечение двух неупорядоченных множеств. Решение с использованием мапы.
// Временная сложность: O(n + m), где n длина a1, а m - длина a2
func intersection[T comparable](a1, a2 []T) (result []T) {
	exists := make(map[T]bool, len(a1))

	for _, v := range a1 {
		exists[v] = true
	}

	for _, v := range a2 {
		if exists[v] {
			result = append(result, v)
			delete(exists, v) // удаляем ключ, чтобы избежать повторений в итоговом множестве
		}
	}
	return
}
