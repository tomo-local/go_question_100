package main

import (
	"fmt"
	"time"
)

func producer(ch chan int){
	for i := 0; i < 4; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go producer(ch)

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("Channel closed, consumer done.")
}
