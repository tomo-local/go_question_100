package main

import(
	"fmt"
	"os"
)

func main(){
	args := os.Args

	// mainを実行のargがあるので-1する
	argsCount := len(args)
	fmt.Printf("count: %d\n", argsCount)

	for i, arg := range args[1:] {
		fmt.Printf("%d: %s\n", i+1, arg)
	}
}
