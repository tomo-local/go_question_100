package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// 1. もとの os.Stdin, os.Stdout を保存しておき、最後に元に戻す
	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	// 2. 標準入力の代わりを作る (Pipe)
	rIn, wIn, _ := os.Pipe()
	os.Stdin = rIn

	// 3. 標準出力の代わりを作る (Pipe)
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	// 4. 入力データを流し込む
	input := "Hello from Test"
	go func() {
		defer wIn.Close()
		wIn.Write([]byte(input + "\n"))
	}()

	// 5. mainを実行
	// 注: main内のfmt.PrintなどがwOutへ書き込まれる
	main()

	// 6. 出力を回収する
	wOut.Close()
	var buf bytes.Buffer
	io.Copy(&buf, rOut)
	got := buf.String()

	// 7.期待する出力の定義
	wantLines := []string{
		"何か入力してください "
		">入力された内容: Hello from Test"
	}

	want := strings.Join(wantLines, "\n") + "\n"

	if got != want {
		t.Errorf("got:\n%q\nwant:\n%q", got, want)
	}
}
