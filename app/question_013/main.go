package main

import(
	"fmt"
)

func main()  {
	fruits := make(map[string]int)
	fruits["apple"] = 100
	fruits["banana"] = 200

	fmt.Printf("key apple: %d\n", fruits["apple"])

	for v, i := range fruits {
		fmt.Printf("%s: %d\n", v,  i)
	}
}
