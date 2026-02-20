package main

import "fmt"

type MyInt int

func (mi MyInt) isEven() bool {
	return mi%2 == 0
}


func main() {
	var i MyInt = 2

	fmt.Printf("i.isEven:%t\n", i.isEven())

	var n MyInt = 1

	fmt.Printf("n.isEven:%t\n", n.isEven())
}
