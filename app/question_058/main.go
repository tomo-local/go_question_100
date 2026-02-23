package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		time.Sleep(1000 * time.Microsecond)
		messages <- "Hello!"
	}()

	for {
		select {
		case msg := <-messages:
			fmt.Println("Received:", msg)
			return
		default:
			fmt.Println("no message received")
		}
		time.Sleep(900 * time.Microsecond)
	}
}
