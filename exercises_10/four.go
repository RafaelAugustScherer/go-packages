package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c, c2 := gen(q)

	receive(c, c2, q)

	fmt.Println("about to exit")
}

func receive(c, c2, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case v := <-c2:
			fmt.Println(v)
		case <-q:
			return
		}

	}
}

func gen(q chan<- int) (<-chan int, <-chan int) {
	c := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				c <- i
			} else {
				c2 <- i
			}
		}
		q <- 1
		close(c)
		close(c2)
	}()

	return c, c2
}
