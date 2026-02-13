package main

import(
	"fmt"
	"strconv"
)

func main(){
	str := "123"
	i, err := strconv.Atoi(str)

	if err != nil {
		fmt.Printf("変換に失敗しまいた: %v\n", err)
	} else {
		fmt.Printf("i: %d\n",i)
	}

	num := 456
	s := strconv.Itoa(num)
	fmt.Printf("s: %s\n", s)
}
