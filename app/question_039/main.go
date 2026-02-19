package main

import "fmt"

func main() {
	p := new(int)

	fmt.Printf("p: %d\n", *p)

	*p = 100
	fmt.Printf("p: %d\n", *p)
}
