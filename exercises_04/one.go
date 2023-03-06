package main

import (
	"fmt"
)

func main() {
	myArray := make([]int, 5)

	myArray[0] = 42
	myArray[1] = 24
	myArray[2] = 27
	myArray[3] = 33
	myArray[4] = 35

	for _, v := range myArray {
		fmt.Println(v)
	}

	fmt.Printf("%T", myArray)
}
