package main

import (
	"fmt"
)

func main() {
	myArray := []int{
		42, 24, 27, 33, 35, 1, 7, 60, 55, 54,
	}

	for _, v := range myArray {
		fmt.Println(v)
	}

	fmt.Printf("%T", myArray)
}
