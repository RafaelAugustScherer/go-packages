package main

import (
	"fmt"
)

func main() {
	const (
		a int = 2
		b string = "2"
		c bool = true
	)

	const (
		d = 2
		e = "2"
		f = true
	)

	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}
