package main

import (
	"fmt"
	"math"
)

func main(){
	x := 3.14
	fmt.Printf("x: %f\n", x)

	y := math.Ceil(x)
	fmt.Printf("y Ceil: %f\n", y)

	z := math.Floor(x)
	fmt.Printf("z Floor: %f\n", z)
}
