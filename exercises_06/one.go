package main

import (
	"fmt"
)

func foo() int {
	return 45
}

func bar() (int, string) {
	return 46, "Just a string"
}

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}
