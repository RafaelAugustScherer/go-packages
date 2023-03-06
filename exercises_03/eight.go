package main

import (
	"fmt"
)

func main() {
	switch {
		case 1 > 2:
			fmt.Println("The sun just crashed on earth :(")
		case 2 > 1:
			fmt.Println("Everything seems normal :D")
	}
}
