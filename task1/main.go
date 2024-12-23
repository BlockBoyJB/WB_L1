package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) WhoAmi() {
	fmt.Printf("Hello! My name is %s and I am %d years old\n", h.Name, h.Age)
}

type Action struct {
	Human
}

func main() {
	a := Action{
		Human{
			Name: "Foobar",
			Age:  20,
		},
	}
	a.WhoAmi()
}
