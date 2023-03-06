package main

import (
	"fmt"
)

func main() {
	sum := func(a int, b int) int {
		return a + b
	}

	exec_op := func(
		op func(int, int) int,
		start_value int,
		ints ...int,
	) int {
		result := 0

		for i, v := range ints {
			fmt.Println(i, v, result)
			if i == 0 {
				result = op(start_value, v)
			} else {
				result = op(result, v)
			}
		}

		return result
	}

	fmt.Println(exec_op(sum, 0, 1, 2, 3, 4, 7, 10))
}
