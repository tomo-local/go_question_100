package main

import "fmt"

func main() {

	i := 1

	for i <= 100 {
		i = i * 2
	}

	fmt.Printf(`i: %d`,i)
}
