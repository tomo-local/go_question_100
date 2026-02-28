package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("tempfile.txt")
	if err != nil {
		fmt.Println("Create error: ", err)
		f.Close()
		return
	}

	_, err = f.Write([]byte("Hello, file!"))
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	if err := f.Close(); err != nil {
		fmt.Println("Close error:", err)
		return
	}
	fmt.Println("Created and wrote to tempfile.txt")

	if err := os.Rename("tempfile.txt", "renamed_tempfile.txt"); err != nil {
		fmt.Println("Rename error:", err)
		return
	}
	fmt.Println("Renamed to renamed_tempfile.txt")

	if err := os.Remove("renamed_tempfile.txt"); err != nil {
		fmt.Println("Remove error:", err)
		return
	}
	fmt.Println("Removed renamed_tempfile.txt")
}
