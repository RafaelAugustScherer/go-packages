package main

import (
	"fmt"
)

type person struct {
	first_name    string
	last_name     string
	fav_ice_cream []string
}

func main() {
	protagonist := person{
		first_name:    "Gordon",
		last_name:     "Freeman",
		fav_ice_cream: []string{"Vanilla", "Strawberry"},
	}

	friend := person{
		first_name:    "Barney",
		last_name:     "Calhoun",
		fav_ice_cream: []string{"Chocolate", "Mint"},
	}

	charsMap := map[string]person{
		protagonist.last_name: protagonist,
		friend.last_name:      friend,
	}

	for _, p := range charsMap {
		fmt.Println("Character Specs:")
		fmt.Printf("\t%v %v\n", p.first_name, p.last_name)

		fmt.Printf("\tFavorite Ice creams:\n")
		for _, ice_cream := range p.fav_ice_cream {
			fmt.Printf("\t\t%v\n", ice_cream)
		}
	}

}
