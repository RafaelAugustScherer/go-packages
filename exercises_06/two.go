package main

import (
	"fmt"
)

func foo(lint ...int) int {
	sum := 0
	for _, v := range lint {
		sum += v
	}

	return sum
}

func bar(lint []int) int {
	sum := 0
	for _, v := range lint {
		sum += v
	}

	return sum
}

func main() {
	list_of_ints := []int{
		1, 2, 4, 6, 8, 10,
	}

	fmt.Println(foo(list_of_ints...))
	fmt.Println(bar(list_of_ints))
}
