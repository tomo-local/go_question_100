package main

import(
	"fmt"
	"strings"
)

func main(){
	str := []string{"Go", "is", "awesome"}

	fmt.Printf("str: %s\n", str)

	joinedStr := strings.Join(str, " ")

	fmt.Printf("joined str: %s\n", joinedStr)
}
