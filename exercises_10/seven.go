package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup // define wg

	wg.Add(1)
	go func(j int) {
		c <- j
		wg.Done()
	}(1)

	wg.Wait()

	close(c)

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("Exiting now..")
}
