package main

import(
	"fmt"
	"encoding/json"
)

type Product struct {
	ID int
	Name string
}

func main() {
	product := Product{
		ID: 1,
		Name: "app",
	}

	jsonData, err := json.Marshal(product)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("json: %s\n", string(jsonData))

}
