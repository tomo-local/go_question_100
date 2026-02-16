package main

import (
	"fmt"
)

func main() {
	var hour int
	hour = 14

	// if 12 > hour {
	// 	fmt.Println("Morning")
	// } else if 17 > hour {
	// 	fmt.Println("Afternoon")
	// } else {
	// 	fmt.Println("Evening")
	// }

	switch  {
	case hour < 12:
		fmt.Println("Morning")
	case hour < 17:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Evening")
	}
}
