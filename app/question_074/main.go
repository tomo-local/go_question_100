package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, io.Copy!")
	n, err := io.Copy(os.Stdout, r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "copy error: %v\n", err)
		return
	}
	fmt.Printf("\nCopied %d bytes\n", n)
}
