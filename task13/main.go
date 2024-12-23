package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Println(a, b)

	// Не знаю как назвать этот вариант (вариант "обычный")
	b, a = a, b
	fmt.Println(a, b)

	// Арифметический вариант
	a = a - b
	b = a + b
	a = b - a
	fmt.Println(a, b)

	// XOR операция
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Println(a, b)

	// Через умножение/деление
	a = a * b
	b = a / b
	a = a / b
	fmt.Println(a, b)
}
