package main

import (
	"fmt"
)

func main() {
	my_map := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for map_k, arr := range my_map {
		fmt.Println(map_k)
		for k, v := range arr {
			fmt.Printf("\t%v %v\n", k, v)
		}
	}
}
