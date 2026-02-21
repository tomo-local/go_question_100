package main

import "fmt"

func do(i interface{}) string{
	switch  i.(type) {
	case int:
		return fmt.Sprintf("Integer: %d", i)
	case string:
		return fmt.Sprintf("String: %s", i)
	case bool:
		return fmt.Sprintf("Bool: %t", i)
	default:
		return fmt.Sprintf("Unknown type")
	}
}

func main() {
	fmt.Printf("%s\n", do("hello"))
	fmt.Printf("%s\n", do(124))
	fmt.Printf("%s\n", do(true))
}
