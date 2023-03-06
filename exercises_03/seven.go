package main

import (
	"fmt"
)

func main() {
	sunday := 0
	munday := 1
	today := 1

	if (today == sunday) {
		fmt.Println("Happy Sunday!!")
	} else if (today == munday) {
		fmt.Println("Uhh no, its munday!!")
	} else {
		fmt.Println("It's okay, just some days till sunday!")
	}
}
