package main

import(
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	// 3. 書き込みを終了して元の標準出力に戻す
	w.Close()
	os.Stdout = old

	// 4. 出力を読み取る
	var buf bytes.Buffer
	io.Copy(&buf, r)
	got := buf.String()

	// 5. 期待する出力の定義
	// 注意: fmt.Printlnは改行を含みます
	wantLines := []string{
		"10 + 5 = 15 ",
		"10 - 5 = 5 ",
	}
	want := strings.Join(wantLines, "\n") + "\n"

	// 6. 検証
	if got != want {
		t.Errorf("got:\n%q\nwant:\n%q", got, want)
	}
}
