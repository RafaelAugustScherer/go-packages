package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var answer = 42
	var counter int64

	wg.Add(answer)
	for i := 0; i < answer; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
