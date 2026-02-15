package main

import(
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("num: %d", rand.Intn(100))
}
