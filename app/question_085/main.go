package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "https://www.example.com:8080/path/to/resource?key1=value1&key2=value2#fragment"
	u, err := url.Parse(s)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("Host:", u.Host)
	fmt.Println("Path:", u.Path)
	fmt.Println("RawQuery:", u.RawQuery)
	fmt.Println("Fragment:", u.Fragment)

	q := u.Query()
	fmt.Println("key1:", q.Get("key1"))
}
