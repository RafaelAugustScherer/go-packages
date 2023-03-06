package main

import (
	"fmt"
)

func main() {
	sum_op := func() func(ints ...int) int {
		sum := 0

		return func(ints ...int) int {
			for _, v := range ints {
				sum += v
			}

			return sum
		}
	}

	add_to_sum := sum_op()

	r := add_to_sum(10, 100)
	fmt.Println(r)

	r = add_to_sum(500)
	fmt.Println(r)
}
