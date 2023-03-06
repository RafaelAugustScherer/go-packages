package main

import (
	"fmt"
)

func run(who string) {
	fmt.Printf("%v is running!\n", who)
}

func main() {
	defer run("First line")
	run("Second line")
}
