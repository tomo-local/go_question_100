package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	path := "./hello.txt"

	// 事前にデータを書き込んでおく（既存ファイルがあるケースのシミュレーション）
	initialContent := "Old Content"
	err := os.WriteFile(path, []byte(initialContent), 0666)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// mainを実行
	main()

	// 最終的なファイルの内容が期待通りか確認
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read file after main: %v", err)
	}

	expected := "Hello, Go!\n"
	if string(content) != expected {
		t.Errorf("expected %q, but got %q", expected, string(content))
	}
}
