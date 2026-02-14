package main

import(
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("何か入力してください \n>")

	if scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("入力された内容: %s\n", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
	}
}
