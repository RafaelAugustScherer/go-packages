package main

import (
	"fmt"
)

func main() {
	myMultiDArray := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for _, record := range myMultiDArray {
		for _, data := range record {
			fmt.Println(data)
		}
	}


}
