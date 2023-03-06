package main

import (
	"fmt"
)

func main() {
	answer := 42

	fmt.Printf("%d\n", answer)
	fmt.Printf("%b\n", answer)
	fmt.Printf("%x\n", answer)

	y := 42 << 1

	fmt.Printf("%d\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
}
