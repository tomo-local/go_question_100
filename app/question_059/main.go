package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Microsecond)
	defer ticker.Stop()

	count := 0
	for range ticker.C {
		fmt.Printf("Tick ar %v\n", count)
		count++
		if count >= 5 {
			break
		}
	}
}
