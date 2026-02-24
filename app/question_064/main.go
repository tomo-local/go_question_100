package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func initialize()  {
	fmt.Println("Initialization completed.")
}

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			once.Do(initialize)
		}()
	}
	time.Sleep(100 * time.Millisecond)
}
