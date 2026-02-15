package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Err: %v\n", err)
	}

	defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Res: %s\n", string(byteArray))
}
