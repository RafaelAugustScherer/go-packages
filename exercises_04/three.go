package main

import (
	"fmt"
)

func main() {
	myArray := []int{
		42, 24, 27, 33, 35, 1, 7, 60, 55, 54,
	}

	mySliceOne := myArray[:5]
	mySliceTwo := myArray[5:10]
	mySliceThree := myArray[2:7]
	mySliceFour := myArray[1:6]

	fmt.Println(mySliceOne)
	fmt.Println(mySliceTwo)
	fmt.Println(mySliceThree)
	fmt.Println(mySliceFour)
}
