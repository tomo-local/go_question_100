package main

import (
	"fmt"
	"strings"
)

func main(){
	fruits := "apple,banana,orange";
  arr := strings.Split(fruits, ",")

	for _, f := range arr {
		fmt.Printf("%s\n",f)
	}

	fmt.Printf("fruits length: %d\n", len(arr))
}
