package main

import (
	"fmt"
)

func main() {
	spike := struct {
		size float64
		bark bool
		meow bool
	}{
		size: 1.2,
		bark: true,
	}

	fmt.Println("Size:", spike.size)
	fmt.Println("Does bark?", spike.bark)
	fmt.Println("Does meow?", spike.meow)
}
