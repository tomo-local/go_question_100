package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, you've hit /hello")
	})

	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, you've hit /world")
	})

	http.ListenAndServe(":8090", nil)
}
