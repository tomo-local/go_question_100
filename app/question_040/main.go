package main

import "fmt"

type Person struct {
	Age int
}

func (p *Person) celebrate(){
	p.Age++
}

func (p Person) getAge() int {
	return p.Age
}

func main() {
	person := Person{Age: 27}

	fmt.Printf("age: %d\n",person.getAge())

	person.celebrate()

	fmt.Printf("age: %d\n",person.getAge())
}
