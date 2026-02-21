package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, io.Reader!")

	data, err := io.ReadAll(r)
	if err != nil {
		fmt.Printf("Read Error: %d", err)
		return
	}

	fmt.Println(string(data))
}
