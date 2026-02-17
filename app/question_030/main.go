package main

import(
	"fmt"
)

func main() {
	fmt.Println("start")

	defer fmt.Println("deferred")

	fmt.Println("end")
}
