package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data = map[string]string{"key1":"value1"}
	mu sync.RWMutex
)

func main() {
	for i := 0; i < 3; i++ {
		id := 1
		go func() {
			mu.RLock()
			defer mu.RUnlock()
			v := data["key1"]
			fmt.Printf("Reader %d: %s\n", id, v)
			time.Sleep(1 * time.Second)
		}()
	}

	go func() {
		time.Sleep(500 * time.Millisecond)
		mu.Lock()
		defer mu.Unlock()
		data["key2"] = "value2"
		fmt.Println("Writer: added key2")
	}()

	time.Sleep(3 * time.Second)
}
