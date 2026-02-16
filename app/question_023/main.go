package main

import (
	"fmt"
)

func main() {
	var n int

	n = 2

	switch n {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}
}
