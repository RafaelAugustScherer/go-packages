package main

import (
	"fmt"
)

func main() {
	public := func() func() string {
		return func() string {
			return "I am anonymous"
		}
	}

	fmt.Println(public()())
}
