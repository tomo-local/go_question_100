package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)

	go func() {
		channel <- "Hello from channel!"
	}()

	msg := <-channel
	fmt.Println(msg)
}
