package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var wg sync.WaitGroup

func main() {
	var answer = 42
	var counter = 0

	wg.Add(answer)
	for i := 0; i < answer; i++ {
		go func() {
			mu.Lock()
			v := counter
			v += 1
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
