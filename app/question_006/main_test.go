package main

import(
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
	"fmt"
)

func Test_main(t *testing.T){
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

	now := time.Now()
	layout := "2006年01月02日"
	oneLine := fmt.Sprintf("now time: %s", now.Format(layout))

	// 5. 期待する出力の定義
	// 注意: fmt.Printlnは改行を含みます
	wantLines := []string{
		oneLine,
		"t time: 1998年06月24日 00時06分00秒",
	}
	want := strings.Join(wantLines, "\n") + "\n"

	// 6. 検証
	if got != want {
		t.Errorf("got:\n%q\nwant:\n%q", got, want)
	}
}
