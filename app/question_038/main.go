package main

import "fmt"

func double(a *int){
	*a = *a *2
}

func main() {
	n := 5
	fmt.Printf("n: %d\n", n)

	double(&n)

	fmt.Printf("n: %d\n", n)
}
