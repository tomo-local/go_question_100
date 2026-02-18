package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func calc(a int, b int, f func(int, int)int) (result int) {
	result = f(a, b)
	return
}

func main() {
	sum := calc(10, 5, add)
	fmt.Printf("10 + 5 = %d \n", sum)

	diff := calc(10, 5, subtract)
	fmt.Printf("10 - 5 = %d \n", diff)
}
