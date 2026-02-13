package main

import(
	"fmt"
	"sort"
)

func main(){
	nums := []int{10, 2, 8, 5, 3}

	fmt.Printf("nums: %d\n", nums)

	sort.Ints(nums)

	fmt.Printf("sorted nums: %d\n", nums)
}
