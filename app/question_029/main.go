package main

import (
	"fmt"
)

func main() {
	for i, v := range map[string]int{"one": 1, "two": 2} {
		fmt.Printf("Key: %s, Value: %d\n", i, v)
	}
}
