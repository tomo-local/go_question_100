package main

import(
	"os"
	"fmt"
)

func main(){
	path := "./hello.txt"
	oldFile, err := os.ReadFile(path)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	if string(oldFile) != "" {
		fmt.Printf("oldFile: %s\n", string(oldFile))
		os.WriteFile(path, []byte(""), 0666)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("file: %s\n", string(file))

	os.WriteFile(path, []byte("Hello, Go!\n"), 0666)

	newFile, err := os.ReadFile(path)

	fmt.Printf("new file: %s\n", string(newFile))
}
