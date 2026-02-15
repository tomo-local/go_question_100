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
	pStr := []byte(`{"id":101,"name":"Book"}`)

	var a Product

	err := json.Unmarshal(pStr, &a)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	fmt.Printf("Success: %+v\n", a)
	fmt.Printf("ID: %d\n", a.ID)
	fmt.Printf("Name: %s\n", a.Name)
}
