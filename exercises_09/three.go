package main

import (
	"fmt"
	"runtime"
	"sync"
)

// go run three.go
// go run -race three.go

var wg sync.WaitGroup

func main() {
	var answer = 42
	var counter = 0

	wg.Add(answer)
	for i := 0; i < answer; i++ {
		go func() {
			v := counter

			runtime.Gosched()
			v += 1
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
