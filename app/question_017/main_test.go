package main

import(
	"bytes"
	"io"
	"os"
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
	want := `Res: {
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}
`

	// 6. 検証
	if got != want {
		t.Errorf("got:\n%q\nwant:\n%q", got, want)
	}
}
