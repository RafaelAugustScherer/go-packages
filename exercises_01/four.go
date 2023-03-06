package main

import (
	"fmt"
)

type MyOwnType int

var x MyOwnType

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
