package main

import (
	"fmt"
)

func isTrue() bool {
	fmt.Println("isTrue() called")
	return true
}

func isFalse() bool {
	fmt.Println("isFalse() called")
	return false
}

func main() {
	if isFalse() && isTrue() {}

	if isTrue() || isFalse() {}
}
