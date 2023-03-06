package main

import (
	"fmt"
)

func main() {
	x := 1

	for {
		fmt.Println(2000 + x)

		if x == 22 {
			break
		}
		x++
	}
}
