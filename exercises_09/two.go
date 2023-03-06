package main

import "fmt"

type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Printf("My name is %v, %v %v\n", p.first, p.first, p.last)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	mr_james := person{
		"James", "Bond",
	}

	saySomething(&mr_james)
	saySomething(mr_james) // WILL FAIL
}
