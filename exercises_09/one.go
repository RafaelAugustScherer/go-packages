package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	fmt.Println("I am foo!")
	wg.Done()
}

func bar() {
	fmt.Println("My name is bar.")
	wg.Done()
}

func main() {
	wg.Add(2)

	go foo()
	go bar()

	wg.Wait()
}
