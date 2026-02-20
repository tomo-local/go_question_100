package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("value: %v type: %T\n", i, i)
}

func main() {
	describe(24)
	describe("tomo")
	describe(true)
	describe(3.14)
}
