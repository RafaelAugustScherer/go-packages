package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(
		"My name is", p.first, p.last,
		"and I`m", p.age, "years old",
	)
}

func main() {
	p1 := person{
		first: "Andre",
		last:  "Hawk",
		age:   32,
	}

	p1.speak()
}
