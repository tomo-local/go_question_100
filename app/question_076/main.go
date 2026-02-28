package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "CUSTOM: ", log.Lshortfile)
	logger.Println("This is a custom log message.")
}
