package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02"
	str := "2023-10-27"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}
	fmt.Println("Parsed:", t)

	formatted := t.Format("Monday, Jan 2, 2006")
	fmt.Println("Formatted:", formatted)
}
