package main

import (
	"fmt"
)

func main() {
	const (
		a = 2024 + iota
		b = 2024 + iota
		c = 2024 + iota
		d = 2024 + iota
	)

	fmt.Println(a, b, c, d)
}
