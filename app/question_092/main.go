package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func ExampleAdd() {
	fmt.Println(add(2, 3))
	// Output:
	// 5
}

func main() {}
