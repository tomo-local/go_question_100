package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf(`Name: %s, Age: %d`, p.Name, p.Age)
}

func main() {
	tomo := Person{Name: "tomo", Age: 27}

	fmt.Println(tomo)
}
