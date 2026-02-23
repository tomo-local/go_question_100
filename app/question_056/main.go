package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Channel closed, range loop ended.")
}
