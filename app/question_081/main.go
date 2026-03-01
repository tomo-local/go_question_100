package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	in := `first_name,last_name,username
Rob,Pike,rob
Ken,Thompson,ken
`

	r := csv.NewReader(strings.NewReader(in))
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("ReadAll error:", err)
		return
	}

	for _, record := range records {
		fmt.Println(record)
	}
}
