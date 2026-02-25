package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker cancelled")
				time.Sleep(1 * time.Second)
			default:
				fmt.Println("working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
