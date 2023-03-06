package main

import (
	"fmt"
)

func main() {
	answer := 42
	rewsna := 24

	g:= answer == rewsna
	h:= answer <= rewsna
	i:= answer >= rewsna
	j:= answer != rewsna
	k:= answer > rewsna
	l:= answer < rewsna

	fmt.Println(answer, rewsna)
	fmt.Printf("Equal: %t\n", g)
	fmt.Printf("Less than or Equal: %t\n", h)
	fmt.Printf("Greater than or Equal: %t\n", i)
	fmt.Printf("Not Equal: %t\n", j)
	fmt.Printf("Greater than: %t\n", k)
	fmt.Printf("Less than: %t\n", l)
}
