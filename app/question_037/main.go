package main

import "fmt"

func main() {
	x := 10

	fmt.Printf("x: %d\n", x)

	p := &x
	fmt.Printf("p: %d\n", *p)

	*p = 20
	fmt.Printf("x: %d\n", x)
}
