package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	select {
	case <-done:
		fmt.Println("Operation completed.")
	case <-ctx.Done():
		fmt.Println("Operation timed out!")
	}

}
