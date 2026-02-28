package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := filepath.Join("dir1", "dir2", "file.txt")
	fmt.Println("Joined path:", path)

	dir := filepath.Dir(path)
	fmt.Println("Dir:", dir)

	base := filepath.Base(path)
	fmt.Println("Base:", base)

	ext := filepath.Ext(path)
	fmt.Println("Ext:", ext)
}
