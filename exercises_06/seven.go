package main

import (
	"fmt"
)

func main() {
	public := func() string {
		return "I am anonymous"
	}

	fmt.Println(public())
}
