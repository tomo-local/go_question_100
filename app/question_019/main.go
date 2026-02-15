package main

import(
	"fmt"
	"strings"
)

func main() {
	str := "Go programming language"

	if strings.Contains(str, "pro") {
		fmt.Printf("contain: %s\n", str)
	}

	if strings.HasPrefix(str, "Go") {
		fmt.Printf("has prefix: %s\n", str)
	}

}
