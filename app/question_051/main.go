package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("Hello from goroutine!")
	} ()

	fmt.Println("Hello from main function!")

	// goroutine の表示が完了するまで待機
	time.Sleep(1 * time.Second)
}
