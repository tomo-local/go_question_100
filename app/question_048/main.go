package main

import(
	"fmt"
	"errors"
)

func validateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}

	return nil
}

func main() {
	if err := validateAge(20); err != nil {
		fmt.Printf("Error: %v\n",err)
	} else {
		fmt.Println("OK")
	}

	if err := validateAge(-5); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("OK")
	}
}
