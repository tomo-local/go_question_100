package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("MY_APP_VAR", "1234")
	fmt.Println("MY_APP_VAR: ", os.Getenv("MY_APP_VAR"))

	val := os.Getenv("NON_EXISTENT_VAR")
	fmt.Printf("NON_EXISTENT_VAR (empty? %v): %q\n", val == "", val)

	if v, ok := os.LookupEnv("MY_APP_VAR"); ok {
		fmt.Println("LookupEnv MY_APP_VAR:", v)
	}
	if _, ok := os.LookupEnv("NON_EXISTENT_VAR"); !ok {
		fmt.Println("NON_EXISTENT_VAR is not set")
	}
}
