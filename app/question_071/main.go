package main

import (
	"flag"
	"fmt"
)

func main() {
	word := flag.String("word", "foo", "a string")
	numb := flag.Int("numb", 42, "an int")
	fork := flag.Bool("fork", false, "a bool")

	flag.Parse()

	fmt.Println("word:", *word)
	fmt.Println("numb:", *numb)
	fmt.Println("fork:", *fork)
	fmt.Println("tail args:", flag.Args())
}
