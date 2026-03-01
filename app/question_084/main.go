package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-a", "-l", "-h")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("exec error:", err)
		return
	}
	fmt.Println(string(out))
}
