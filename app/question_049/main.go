package main

import (
	"fmt"
)

type MyError struct {
	Code int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error: %d: %s", e.Code, e.Message)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &MyError{
			Code:    1001,
			Message: "division by zero",
		}
	}
	return a / b, nil
}

func main() {
	a, b := 10, 0
	result, err := divide(a, b)

	if err != nil {
		fmt.Printf("divide Error: %v\n", err)

		if myErr, ok := err.(*MyError); ok {
			fmt.Println("Custom Error")
			fmt.Printf("Error Code: %d\n", myErr.Code)
			fmt.Printf("Error Message: %s\n", myErr.Message)
		}
	} else {
		fmt.Printf("%d ÷ %d = %d\n", a, b, result)
	}
}
