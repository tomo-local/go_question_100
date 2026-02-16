package main

import (
	"fmt"
)

func main() {
	var score int

	score = 85

	if score >= 80 {
		fmt.Println("Great")
	} else if score >= 60 {
		fmt.Println("Good")
	} else {
		fmt.Println("OK")
	}

}
