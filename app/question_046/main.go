package main

import (
	"fmt"
	"strings"
)

func main() {
	var i interface{} = "hello"

	s, ok := i.(string)

	if ok {
		fmt.Println("i.(string):", strings.ToUpper(s))
	}

	var i2 interface{} = 123

	s2, ok := i2.(int)

	if ok {
		fmt.Printf("i.(int): %d\n", s2)
	}
}
