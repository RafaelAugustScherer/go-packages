package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	p.first = "Larry"
}

func main() {
	p1 := person{
		first: "Homer",
		last:  "Simpson",
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
