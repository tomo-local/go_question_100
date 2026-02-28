package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	f, err := os.Create("users.csv")
	if err != nil {
		fmt.Println("Create error:", err)
		return
	}

	// 処理が完了したら、ファイルを閉じて削除する
	defer func() {
		f.Close()
		defer os.Remove("users.csv")
		fmt.Println("Remove users.csv")
	}()

	w := csv.NewWriter(f)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			fmt.Println("Write error:", err)
			return
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Println("Writer error:", err)
		return
	}
	fmt.Println("Wrote users.csv")

}
