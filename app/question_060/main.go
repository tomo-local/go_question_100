package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)

	select {
	case <- timer.C:
		fmt.Println("Timer expired!")
	}
}
