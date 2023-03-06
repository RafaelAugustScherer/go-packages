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

	my_map["adam_smith"] = []string{`Cool Guy`, `Book`, `Math`}

	delete(my_map, "no_dr")

	for map_k, arr := range my_map {
		fmt.Printf("%v\n\t%v\n", map_k, arr)
	}
}
