package main

import (
	"fmt"
)

func main(){
	s := []int{1,2,3}

	fmt.Printf("s: %d\n", s)

	s = append(s, 4)

	fmt.Printf("s: %d\n", s)
}
