package main

import (
	"fmt"
)

func main() {
	for i ,v := range []string{"a", "b", "c"} {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}
}
