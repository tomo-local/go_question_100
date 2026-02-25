package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, err := regexp.Compile(`a `)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	s1 := re.ReplaceAllString("a peach", "an ")
	s2 := re.ReplaceAllString("a banana", "an ")

	fmt.Println(s1)
	fmt.Println(s2)
}
