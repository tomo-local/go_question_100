package main

import(
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main() {
	person := Person{
		Name: "tomo",
		Age: 27,
	}

	fmt.Printf("name: %s age: %d\n", person.Name, person.Age)
}
