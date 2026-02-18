package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {

	total := sum(1,2,3,4,5)
	fmt.Printf("1+2+3+4+5=%d\n", total)

	slice := []int{100,200}
	total2 := sum(slice...)
	fmt.Printf("[]int{100,200}=%d\n", total2)
}
