package main

import (
	"fmt"
)

func pinger(pings chan<- string, msg string) {
	pings <- msg
}

func ponger(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	go pinger(pings, "ping")
	go ponger(pings, pongs)

	fmt.Println(<-pongs)

}
