package main

import (
	"fmt"
)

func divmod(a, b int)(quotient, remainder int)  {
	quotient = a/b
	remainder = a%b
	return
}


func main() {
	q, r := divmod(10, 3)
	fmt.Printf("10/3=%d, 余り: %d\n", q, r)
}
