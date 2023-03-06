package main

import (
	"fmt"
)

func main() {
	favSport := "mountain biking"

	switch favSport {
		case "football":
			fmt.Println("Great choice!")
		case "flipping":
			fmt.Println("That's not even a sport")
		case "mountain biking":
			fmt.Println("That's my favorite too!")
	}
}
