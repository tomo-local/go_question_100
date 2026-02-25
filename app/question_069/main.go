package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := `p([a-z]+)ch`
	matched, err := regexp.MatchString(pattern, "peach")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Matched:", matched)

	matched2, err := regexp.MatchString(pattern, "1234")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Matched:", matched2)
}
