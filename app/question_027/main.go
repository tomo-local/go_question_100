package main

import(
	"fmt"
)

func main() {

	i := 0
	for {
		if i == 5 {
			break;
		}
		i++
	}

	fmt.Printf("Loop finished: %d\n", i)
}
