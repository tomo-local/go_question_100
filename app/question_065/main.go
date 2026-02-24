package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var ops uint64

func main() {
	var wg sync.WaitGroup
	const numGoroutines = 50
	const incPerGoroutine = 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incPerGoroutine; j++ {
				atomic.AddUint64(&ops, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("ops =", ops)
}
